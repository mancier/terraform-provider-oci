// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v34/database"
)

func init() {
	RegisterDatasource("oci_database_external_non_container_database", DatabaseExternalNonContainerDatabaseDataSource())
}

func DatabaseExternalNonContainerDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_non_container_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseExternalNonContainerDatabaseResource(), fieldMap, readSingularDatabaseExternalNonContainerDatabase)
}

func readSingularDatabaseExternalNonContainerDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseExternalNonContainerDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExternalNonContainerDatabaseResponse
}

func (s *DatabaseExternalNonContainerDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalNonContainerDatabaseDataSourceCrud) Get() error {
	request := oci_database.GetExternalNonContainerDatabaseRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetExternalNonContainerDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{DatabaseManagementConfigToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DbId != nil {
		s.D.Set("db_id", *s.Res.DbId)
	}

	if s.Res.DbPacks != nil {
		s.D.Set("db_packs", *s.Res.DbPacks)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}
