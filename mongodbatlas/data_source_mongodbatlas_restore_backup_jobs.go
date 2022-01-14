package mongodbatlas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceMongoDBAtlasRestoreBackupJobs() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceMongoDBAtlasRestoreBackupJobsRead,
		Importer: &schema.ResourceImporter{
			StateContext: datasourceMongoDBAtlasRestoreBackupJobsImportState,
		},
		Schema: returnDataSourceRestoreBackupJobsSchema(),
	}
}

func returnDataSourceRestoreBackupJobsSchema() map[string]*schema.Schema {

}

func datasourceMongoDBAtlasRestoreBackupJobsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

}

func datasourceMongoDBAtlasRestoreBackupJobsImportState(ctx context.Context, data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {

}
