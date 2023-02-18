package elasticsearch_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/opensearchservice"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func TestAccElasticSearchEngineVersionDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_elasticsearch_engine_version.test"
	version := "1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccEngineVersionPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, opensearchservice.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineVersionDataSourceConfig_basic(version),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "version", version),
				),
			},
		},
	})

}

func testAccEngineVersionDataSourceConfig_basic(version string) string {
	return fmt.Sprintf(`
data "aws_elasticsearch_engine_version" "test" {
  version                = %[1]q
}
`, version)
}

func testAccEngineVersionPreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).OpenSearchConn()

	input := &opensearchservice.ListVersionsInput{}

	_, err := conn.ListVersionsWithContext(ctx, input)

	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}
