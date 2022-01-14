package mongodbatlas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceMongoDBAtlasRestoreBackupJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMongoDBAtlasRestoreBackupJobCreate,
		ReadContext:   resourceMongoDBAtlasRestoreBackupJobRead,
		UpdateContext: resourceMongoDBAtlasRestoreBackupJobUpdate,
		DeleteContext: resourceMongoDBAtlasRestoreBackupJobDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMongoDBAtlasRestoreBackupJobImportState,
		},
		Schema: returnRestoreBackupJobSchema(),
	}
}

func returnRestoreBackupJobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"project_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cluster_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"delivery_type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"automated", "download", "pointInTime"}, false),
		},
		"snapshot_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"target_cluster_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"target_project_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cancelled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"expired": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"expires_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"finished_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"restore_job_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"links": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"href": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"rel": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"timestamp": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func resourceMongoDBAtlasRestoreBackupJobCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

}

func resourceMongoDBAtlasRestoreBackupJobRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

}

func resourceMongoDBAtlasRestoreBackupJobUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

}

func resourceMongoDBAtlasRestoreBackupJobDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

}

func resourceMongoDBAtlasRestoreBackupJobImportState(ctx context.Context, data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {

}
