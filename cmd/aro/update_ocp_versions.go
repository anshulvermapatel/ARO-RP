package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/metrics/statsd"
	"github.com/Azure/ARO-RP/pkg/util/encryption"
	"github.com/Azure/ARO-RP/pkg/util/keyvault"
)

func getLatestOCPVersions(ctx context.Context, log *logrus.Entry) ([]api.OpenShiftVersion, error) {
	env, err := env.NewCoreForCI(ctx, log)
	if err != nil {
		return nil, err
	}
	dstAcr := os.Getenv("DST_ACR_NAME")
	acrDomainSuffix := "." + env.Environment().ContainerRegistryDNSSuffix

	dstRepo := dstAcr + acrDomainSuffix
	var (
		OpenshiftVersions = []api.OpenShiftVersion{
			{
				Version:           "4.10.20",
				OpenShiftPullspec: "quay.io/openshift-release-dev/ocp-release@sha256:e1fa1f513068082d97d78be643c369398b0e6820afab708d26acda2262940954",
				InstallerPullspec: dstRepo + "/aro-installer:release-4.10",
				Enabled:           true,
			},
		}
	)
	return OpenshiftVersions, nil
}

func getVersionsDatabase(ctx context.Context, log *logrus.Entry) (database.OpenShiftVersions, error) {
	_env, err := env.NewCore(ctx, log)
	if err != nil {
		return nil, err
	}

	for _, key := range []string{
		"DST_ACR_NAME",
	} {
		if _, found := os.LookupEnv(key); !found {
			return nil, fmt.Errorf("environment variable %q unset", key)
		}
	}

	if !_env.IsLocalDevelopmentMode() {
		for _, key := range []string{
			"MDM_ACCOUNT",
			"MDM_NAMESPACE",
		} {
			if _, found := os.LookupEnv(key); !found {
				return nil, fmt.Errorf("environment variable %q unset", key)
			}
		}
	}

	msiAuthorizer, err := _env.NewMSIAuthorizer(env.MSIContextRP, _env.Environment().ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	msiKVAuthorizer, err := _env.NewMSIAuthorizer(env.MSIContextRP, _env.Environment().ResourceIdentifiers.KeyVault)
	if err != nil {
		return nil, err
	}

	m := statsd.New(ctx, log.WithField("component", "update-ocp-versions"), _env, os.Getenv("MDM_ACCOUNT"), os.Getenv("MDM_NAMESPACE"), os.Getenv("MDM_STATSD_SOCKET"))

	serviceKeyvaultURI, err := keyvault.URI(_env, env.ServiceKeyvaultSuffix)
	if err != nil {
		return nil, err
	}

	serviceKeyvault := keyvault.NewManager(msiKVAuthorizer, serviceKeyvaultURI)

	aead, err := encryption.NewMulti(ctx, serviceKeyvault, env.EncryptionSecretV2Name, env.EncryptionSecretName)
	if err != nil {
		return nil, err
	}

	dbAuthorizer, err := database.NewMasterKeyAuthorizer(ctx, _env, msiAuthorizer)
	if err != nil {
		return nil, err
	}

	dbc, err := database.NewDatabaseClient(log.WithField("component", "database"), _env, dbAuthorizer, m, aead)
	if err != nil {
		return nil, err
	}

	dbOpenShiftVersions, err := database.NewOpenShiftVersions(ctx, _env.IsLocalDevelopmentMode(), dbc)
	if err != nil {
		return nil, err
	}

	return dbOpenShiftVersions, nil
}

func updateOpenShiftVersions(ctx context.Context, dbOpenShiftVersions database.OpenShiftVersions, log *logrus.Entry) error {
	existingVersions, err := dbOpenShiftVersions.ListAll(ctx)
	if err != nil {
		return err
	}

	latestVersions, err := getLatestOCPVersions(ctx, log)
	if err != nil {
		return err
	}

	newVersions := make(map[string]api.OpenShiftVersion)
	for _, doc := range latestVersions {
		newVersions[doc.Version] = doc
	}

	for _, doc := range existingVersions.OpenShiftVersionDocuments {
		existing, found := newVersions[doc.OpenShiftVersion.Version]
		if found {
			log.Printf("Found Version %q, patching", existing.Version)
			_, err := dbOpenShiftVersions.Patch(ctx, doc.ID, func(inFlightDoc *api.OpenShiftVersionDocument) error {
				inFlightDoc.OpenShiftVersion = &existing
				return nil
			})
			if err != nil {
				return err
			}
			log.Printf("Version %q found", existing.Version)
			delete(newVersions, existing.Version)
			continue
		}

		log.Printf("Version %q not found, deleting", doc.OpenShiftVersion.Version)
		err := dbOpenShiftVersions.Delete(ctx, doc)
		if err != nil {
			return err
		}
	}

	for _, doc := range newVersions {
		log.Printf("Version %q not found in database, creating", doc.Version)
		newDoc := api.OpenShiftVersionDocument{
			ID:               dbOpenShiftVersions.NewUUID(),
			OpenShiftVersion: &doc,
		}
		_, err := dbOpenShiftVersions.Create(ctx, &newDoc)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateOCPVersions(ctx context.Context, log *logrus.Entry) error {
	dbOpenShiftVersions, err := getVersionsDatabase(ctx, log)
	if err != nil {
		return err
	}

	err = updateOpenShiftVersions(ctx, dbOpenShiftVersions, log)
	if err != nil {
		return err
	}
	return nil
}
