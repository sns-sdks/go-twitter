package twitter

/*
	Compliance include api for compliance
*/

type ComplianceResource Resource

func newComplianceResource(cli *Client) *ComplianceResource {
	return &ComplianceResource{Cli: cli}
}

// ComplianceJob represents the data for compliance job.
// Refer: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs-id
type ComplianceJob struct {
	ID                *string `json:"id"`
	CreatedAt         *string `json:"created_at,omitempty"`
	Type              *string `json:"type,omitempty"`
	Name              *string `json:"name,omitempty"`
	Status            *string `json:"status,omitempty"`
	Resumable         *bool   `json:"resumable,omitempty"`
	UploadURL         *string `json:"upload_url,omitempty"`
	UploadExpiresAt   *string `json:"upload_expires_at,omitempty"`
	DownloadURL       *string `json:"download_url,omitempty"`
	DownloadExpiresAt *string `json:"download_expires_at,omitempty"`
}

func (j ComplianceJob) String() string {
	return Stringify(j)
}

// LookupJobByID Get a single compliance job with the specified ID.
// Refer: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs-id
func (r *ComplianceResource) LookupJobByID(id string) (*ComplianceJobResp, *APIError) {
	path := "/compliance/jobs/" + id

	resp := new(ComplianceJobResp)
	err := r.Cli.DoGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LookupJobsOpts specifies the parameters for get recent jobs.
type LookupJobsOpts struct {
	Type   string `url:"type"`
	Status string `url:"status,omitempty"`
}

// LookupJobs Returns a list of recent compliance jobs.
// Refer: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs
func (r *ComplianceResource) LookupJobs(args LookupJobsOpts) (*ComplianceJobsResp, *APIError) {
	path := "/compliance/jobs"

	resp := new(ComplianceJobsResp)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateJobOpts specifies the parameters for create a compliance job.
type CreateJobOpts struct {
	Type      string `url:"type"`
	Name      string `url:"name,omitempty"`
	Resumable bool   `url:"resumable,omitempty"`
}

// CreateJob Creates a new compliance job for Tweet IDs or user IDs.
// Refer: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/post-compliance-jobs
func (r *ComplianceResource) CreateJob(args CreateJobOpts) (*ComplianceJobResp, *APIError) {
	path := "/compliance/jobs"

	resp := new(ComplianceJobResp)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
