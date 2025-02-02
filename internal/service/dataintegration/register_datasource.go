// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_dataintegration_workspace", DataintegrationWorkspaceDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_application", DataintegrationWorkspaceApplicationDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_application_patch", DataintegrationWorkspaceApplicationPatchDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_application_patches", DataintegrationWorkspaceApplicationPatchesDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_application_schedule", DataintegrationWorkspaceApplicationScheduleDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_application_schedules", DataintegrationWorkspaceApplicationSchedulesDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_applications", DataintegrationWorkspaceApplicationsDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_export_request", DataintegrationWorkspaceExportRequestDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_export_requests", DataintegrationWorkspaceExportRequestsDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_folder", DataintegrationWorkspaceFolderDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_folders", DataintegrationWorkspaceFoldersDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_import_request", DataintegrationWorkspaceImportRequestDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_import_requests", DataintegrationWorkspaceImportRequestsDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_project", DataintegrationWorkspaceProjectDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspace_projects", DataintegrationWorkspaceProjectsDataSource())
	tfresource.RegisterDatasource("oci_dataintegration_workspaces", DataintegrationWorkspacesDataSource())
}
