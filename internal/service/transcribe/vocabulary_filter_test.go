package transcribe_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftranscribe "github.com/hashicorp/terraform-provider-aws/internal/service/transcribe"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccTranscribeVocabularyFilter_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var vocabularyFilter transcribe.GetVocabularyFilterOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_transcribe_vocabulary_filter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(names.TranscribeEndpointID, t)
			testAccVocabularyFiltersPreCheck(t)
		},
		ErrorCheck:        acctest.ErrorCheck(t, names.TranscribeEndpointID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckVocabularyFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyFilterConfig_basicFile(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttrSet(resourceName, "arn"),
					resource.TestCheckResourceAttrSet(resourceName, "download_uri"),
					resource.TestCheckResourceAttr(resourceName, "language_code", "en-US"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vocabulary_filter_file_uri", "download_uri"},
			},
		},
	})
}

func TestAccTranscribeVocabularyFilter_basicWords(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var vocabularyFilter transcribe.GetVocabularyFilterOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_transcribe_vocabulary_filter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(names.TranscribeEndpointID, t)
			testAccVocabularyFiltersPreCheck(t)
		},
		ErrorCheck:        acctest.ErrorCheck(t, names.TranscribeEndpointID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckVocabularyFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyFilterConfig_basicWords(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttrSet(resourceName, "arn"),
					resource.TestCheckResourceAttrSet(resourceName, "download_uri"),
					resource.TestCheckResourceAttr(resourceName, "language_code", "en-US"),
				),
			},
		},
	})
}

func TestAccTranscribeVocabularyFilter_update(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var vocabularyFilter transcribe.GetVocabularyFilterOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_transcribe_vocabulary_filter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(names.TranscribeEndpointID, t)
			testAccVocabularyFiltersPreCheck(t)
		},
		ErrorCheck:        acctest.ErrorCheck(t, names.TranscribeEndpointID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckVocabularyFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyFilterConfig_basicFile(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttrSet(resourceName, "arn"),
					resource.TestCheckResourceAttrSet(resourceName, "download_uri"),
					resource.TestCheckResourceAttr(resourceName, "vocabulary_filter_file_uri", "s3://"+rName+"/transcribe/test1.txt"),
				),
			},
			{
				Config: testAccVocabularyFilterConfig_basicWords(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttrSet(resourceName, "arn"),
					resource.TestCheckResourceAttrSet(resourceName, "download_uri"),
					resource.TestCheckResourceAttr(resourceName, "words.#", "3"),
				),
			},
		},
	})
}

func TestAccTranscribeVocabularyFilter_updateTags(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var vocabularyFilter transcribe.GetVocabularyFilterOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_transcribe_vocabulary_filter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(names.TranscribeEndpointID, t)
			testAccVocabularyFiltersPreCheck(t)
		},
		ErrorCheck:        acctest.ErrorCheck(t, names.TranscribeEndpointID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckVocabularyFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyFilterConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				Config: testAccVocabularyFilterConfig_tags2(rName, "key1", "value1", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccVocabularyFilterConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccTranscribeVocabularyFilter_disappears(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var vocabularyFilter transcribe.GetVocabularyFilterOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_transcribe_vocabulary_filter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(names.TranscribeEndpointID, t)
			testAccVocabularyFiltersPreCheck(t)
		},
		ErrorCheck:        acctest.ErrorCheck(t, names.TranscribeEndpointID),
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccCheckVocabularyFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVocabularyFilterConfig_basicFile(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocabularyFilterExists(resourceName, &vocabularyFilter),
					acctest.CheckResourceDisappears(acctest.Provider, tftranscribe.ResourceVocabularyFilter(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckVocabularyFilterDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).TranscribeConn
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_transcribe_vocabulary_filter" {
			continue
		}

		_, err := tftranscribe.FindVocabularyFilterByName(ctx, conn, rs.Primary.ID)

		if tfresource.NotFound(err) {
			continue
		}

		if err != nil {
			return err
		}

		return names.Error(names.Transcribe, names.ErrActionCheckingDestroyed, tftranscribe.ResNameVocabularyFilter, rs.Primary.ID, errors.New("not destroyed"))
	}

	return nil
}

func testAccCheckVocabularyFilterExists(name string, vocabularyFilter *transcribe.GetVocabularyFilterOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return names.Error(names.Transcribe, names.ErrActionCheckingExistence, tftranscribe.ResNameVocabularyFilter, name, errors.New("not found"))
		}

		if rs.Primary.ID == "" {
			return names.Error(names.Transcribe, names.ErrActionCheckingExistence, tftranscribe.ResNameVocabularyFilter, name, errors.New("not set"))
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).TranscribeConn
		ctx := context.Background()
		resp, err := tftranscribe.FindVocabularyFilterByName(ctx, conn, rs.Primary.ID)

		if err != nil {
			return names.Error(names.Transcribe, names.ErrActionCheckingExistence, tftranscribe.ResNameVocabularyFilter, rs.Primary.ID, err)
		}

		*vocabularyFilter = *resp

		return nil
	}
}

func testAccVocabularyFiltersPreCheck(t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).TranscribeConn
	ctx := context.Background()

	input := &transcribe.ListVocabularyFiltersInput{}
	_, err := conn.ListVocabularyFilters(ctx, input)

	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}

func testAccVocabularyFilterBaseConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "test" {
  bucket        = %[1]q
  force_destroy = true
}

resource "aws_s3_object" "object1" {
  bucket = aws_s3_bucket.test.id
  key    = "transcribe/test1.txt"
  source = "test-fixtures/vocabulary_filter_test1.txt"
}
`, rName)
}

func testAccVocabularyFilterConfig_basicFile(rName string) string {
	return acctest.ConfigCompose(
		testAccVocabularyFilterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_transcribe_vocabulary_filter" "test" {
  vocabulary_filter_name     = %[1]q
  language_code              = "en-US"
  vocabulary_filter_file_uri = "s3://${aws_s3_bucket.test.id}/${aws_s3_object.object1.key}"

  tags = {
    tag1 = "value1"
    tag2 = "value3"
  }

  depends_on = [
    aws_s3_object.object1
  ]
}
`, rName))
}

func testAccVocabularyFilterConfig_basicWords(rName string) string {
	return acctest.ConfigCompose(
		testAccVocabularyBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_transcribe_vocabulary_filter" "test" {
  vocabulary_filter_name = %[1]q
  language_code          = "en-US"
  words                  = ["bucket", "cars", "boards"]

  tags = {
    tag1 = "value1"
    tag2 = "value3"
  }
}
`, rName))
}

func testAccVocabularyFilterConfig_tags1(rName, key1, value1 string) string {
	return acctest.ConfigCompose(
		testAccVocabularyFilterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_transcribe_vocabulary_filter" "test" {
  vocabulary_filter_name     = %[1]q
  language_code              = "en-US"
  vocabulary_filter_file_uri = "s3://${aws_s3_bucket.test.id}/${aws_s3_object.object1.key}"

  tags = {
    %[2]q = %[3]q
  }

  depends_on = [
    aws_s3_object.object1
  ]
}
`, rName, key1, value1))
}

func testAccVocabularyFilterConfig_tags2(rName, key1, value1, key2, value2 string) string {
	return acctest.ConfigCompose(
		testAccVocabularyFilterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_transcribe_vocabulary_filter" "test" {
  vocabulary_filter_name     = %[1]q
  language_code              = "en-US"
  vocabulary_filter_file_uri = "s3://${aws_s3_bucket.test.id}/${aws_s3_object.object1.key}"

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }

  depends_on = [
    aws_s3_object.object1
  ]
}
`, rName, key1, value1, key2, value2))
}
