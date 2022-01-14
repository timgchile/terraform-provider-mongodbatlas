package mongodbatlas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceMongoDBAtlasRestoreBackupJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceMongoDBAtlasRestoreBackupJobRead,
		Importer: &schema.ResourceImporter{
			StateContext: datasourceMongoDBAtlasRestoreBackupJobImportState,
		},
		Schema: returnDataSourceRestoreBackupJobSchema(),
	}
}

func returnDataSourceRestoreBackupJobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"project_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cluster_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"index_id": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"analyzer": {
			Type:     schema.TypeString,
			Required: true,
		},
		"analyzers": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: validateSearchAnalyzersDiff,
		},
		"collection_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"database": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"search_analyzer": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mappings_dynamic": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"mappings_fields": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: validateSearchIndexMappingDiff,
		},
		"synonyms": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"analyzer": {
						Type:     schema.TypeString,
						Required: true,
					},
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"source_collection": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func datasourceMongoDBAtlasRestoreBackupJobImportState(ctx context.Context, data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {

}

func datasourceMongoDBAtlasRestoreBackupJobRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

}
