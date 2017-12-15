// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lexmodelbuildingserviceiface provides an interface to enable mocking the Amazon Lex Model Building Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lexmodelbuildingserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
)

// LexModelBuildingServiceAPI provides an interface to enable mocking the
// lexmodelbuildingservice.LexModelBuildingService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lex Model Building Service.
//    func myFunc(svc lexmodelbuildingserviceiface.LexModelBuildingServiceAPI) bool {
//        // Make svc.CreateBotVersion request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := lexmodelbuildingservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLexModelBuildingServiceClient struct {
//        lexmodelbuildingserviceiface.LexModelBuildingServiceAPI
//    }
//    func (m *mockLexModelBuildingServiceClient) CreateBotVersion(input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLexModelBuildingServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LexModelBuildingServiceAPI interface {
	CreateBotVersionRequest(*lexmodelbuildingservice.CreateBotVersionInput) lexmodelbuildingservice.CreateBotVersionRequest

	CreateIntentVersionRequest(*lexmodelbuildingservice.CreateIntentVersionInput) lexmodelbuildingservice.CreateIntentVersionRequest

	CreateSlotTypeVersionRequest(*lexmodelbuildingservice.CreateSlotTypeVersionInput) lexmodelbuildingservice.CreateSlotTypeVersionRequest

	DeleteBotRequest(*lexmodelbuildingservice.DeleteBotInput) lexmodelbuildingservice.DeleteBotRequest

	DeleteBotAliasRequest(*lexmodelbuildingservice.DeleteBotAliasInput) lexmodelbuildingservice.DeleteBotAliasRequest

	DeleteBotChannelAssociationRequest(*lexmodelbuildingservice.DeleteBotChannelAssociationInput) lexmodelbuildingservice.DeleteBotChannelAssociationRequest

	DeleteBotVersionRequest(*lexmodelbuildingservice.DeleteBotVersionInput) lexmodelbuildingservice.DeleteBotVersionRequest

	DeleteIntentRequest(*lexmodelbuildingservice.DeleteIntentInput) lexmodelbuildingservice.DeleteIntentRequest

	DeleteIntentVersionRequest(*lexmodelbuildingservice.DeleteIntentVersionInput) lexmodelbuildingservice.DeleteIntentVersionRequest

	DeleteSlotTypeRequest(*lexmodelbuildingservice.DeleteSlotTypeInput) lexmodelbuildingservice.DeleteSlotTypeRequest

	DeleteSlotTypeVersionRequest(*lexmodelbuildingservice.DeleteSlotTypeVersionInput) lexmodelbuildingservice.DeleteSlotTypeVersionRequest

	DeleteUtterancesRequest(*lexmodelbuildingservice.DeleteUtterancesInput) lexmodelbuildingservice.DeleteUtterancesRequest

	GetBotRequest(*lexmodelbuildingservice.GetBotInput) lexmodelbuildingservice.GetBotRequest

	GetBotAliasRequest(*lexmodelbuildingservice.GetBotAliasInput) lexmodelbuildingservice.GetBotAliasRequest

	GetBotAliasesRequest(*lexmodelbuildingservice.GetBotAliasesInput) lexmodelbuildingservice.GetBotAliasesRequest

	GetBotAliasesPages(*lexmodelbuildingservice.GetBotAliasesInput, func(*lexmodelbuildingservice.GetBotAliasesOutput, bool) bool) error
	GetBotAliasesPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotAliasesInput, func(*lexmodelbuildingservice.GetBotAliasesOutput, bool) bool, ...aws.Option) error

	GetBotChannelAssociationRequest(*lexmodelbuildingservice.GetBotChannelAssociationInput) lexmodelbuildingservice.GetBotChannelAssociationRequest

	GetBotChannelAssociationsRequest(*lexmodelbuildingservice.GetBotChannelAssociationsInput) lexmodelbuildingservice.GetBotChannelAssociationsRequest

	GetBotChannelAssociationsPages(*lexmodelbuildingservice.GetBotChannelAssociationsInput, func(*lexmodelbuildingservice.GetBotChannelAssociationsOutput, bool) bool) error
	GetBotChannelAssociationsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotChannelAssociationsInput, func(*lexmodelbuildingservice.GetBotChannelAssociationsOutput, bool) bool, ...aws.Option) error

	GetBotVersionsRequest(*lexmodelbuildingservice.GetBotVersionsInput) lexmodelbuildingservice.GetBotVersionsRequest

	GetBotVersionsPages(*lexmodelbuildingservice.GetBotVersionsInput, func(*lexmodelbuildingservice.GetBotVersionsOutput, bool) bool) error
	GetBotVersionsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotVersionsInput, func(*lexmodelbuildingservice.GetBotVersionsOutput, bool) bool, ...aws.Option) error

	GetBotsRequest(*lexmodelbuildingservice.GetBotsInput) lexmodelbuildingservice.GetBotsRequest

	GetBotsPages(*lexmodelbuildingservice.GetBotsInput, func(*lexmodelbuildingservice.GetBotsOutput, bool) bool) error
	GetBotsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBotsInput, func(*lexmodelbuildingservice.GetBotsOutput, bool) bool, ...aws.Option) error

	GetBuiltinIntentRequest(*lexmodelbuildingservice.GetBuiltinIntentInput) lexmodelbuildingservice.GetBuiltinIntentRequest

	GetBuiltinIntentsRequest(*lexmodelbuildingservice.GetBuiltinIntentsInput) lexmodelbuildingservice.GetBuiltinIntentsRequest

	GetBuiltinIntentsPages(*lexmodelbuildingservice.GetBuiltinIntentsInput, func(*lexmodelbuildingservice.GetBuiltinIntentsOutput, bool) bool) error
	GetBuiltinIntentsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinIntentsInput, func(*lexmodelbuildingservice.GetBuiltinIntentsOutput, bool) bool, ...aws.Option) error

	GetBuiltinSlotTypesRequest(*lexmodelbuildingservice.GetBuiltinSlotTypesInput) lexmodelbuildingservice.GetBuiltinSlotTypesRequest

	GetBuiltinSlotTypesPages(*lexmodelbuildingservice.GetBuiltinSlotTypesInput, func(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, bool) bool) error
	GetBuiltinSlotTypesPagesWithContext(aws.Context, *lexmodelbuildingservice.GetBuiltinSlotTypesInput, func(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, bool) bool, ...aws.Option) error

	GetExportRequest(*lexmodelbuildingservice.GetExportInput) lexmodelbuildingservice.GetExportRequest

	GetIntentRequest(*lexmodelbuildingservice.GetIntentInput) lexmodelbuildingservice.GetIntentRequest

	GetIntentVersionsRequest(*lexmodelbuildingservice.GetIntentVersionsInput) lexmodelbuildingservice.GetIntentVersionsRequest

	GetIntentVersionsPages(*lexmodelbuildingservice.GetIntentVersionsInput, func(*lexmodelbuildingservice.GetIntentVersionsOutput, bool) bool) error
	GetIntentVersionsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetIntentVersionsInput, func(*lexmodelbuildingservice.GetIntentVersionsOutput, bool) bool, ...aws.Option) error

	GetIntentsRequest(*lexmodelbuildingservice.GetIntentsInput) lexmodelbuildingservice.GetIntentsRequest

	GetIntentsPages(*lexmodelbuildingservice.GetIntentsInput, func(*lexmodelbuildingservice.GetIntentsOutput, bool) bool) error
	GetIntentsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetIntentsInput, func(*lexmodelbuildingservice.GetIntentsOutput, bool) bool, ...aws.Option) error

	GetSlotTypeRequest(*lexmodelbuildingservice.GetSlotTypeInput) lexmodelbuildingservice.GetSlotTypeRequest

	GetSlotTypeVersionsRequest(*lexmodelbuildingservice.GetSlotTypeVersionsInput) lexmodelbuildingservice.GetSlotTypeVersionsRequest

	GetSlotTypeVersionsPages(*lexmodelbuildingservice.GetSlotTypeVersionsInput, func(*lexmodelbuildingservice.GetSlotTypeVersionsOutput, bool) bool) error
	GetSlotTypeVersionsPagesWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypeVersionsInput, func(*lexmodelbuildingservice.GetSlotTypeVersionsOutput, bool) bool, ...aws.Option) error

	GetSlotTypesRequest(*lexmodelbuildingservice.GetSlotTypesInput) lexmodelbuildingservice.GetSlotTypesRequest

	GetSlotTypesPages(*lexmodelbuildingservice.GetSlotTypesInput, func(*lexmodelbuildingservice.GetSlotTypesOutput, bool) bool) error
	GetSlotTypesPagesWithContext(aws.Context, *lexmodelbuildingservice.GetSlotTypesInput, func(*lexmodelbuildingservice.GetSlotTypesOutput, bool) bool, ...aws.Option) error

	GetUtterancesViewRequest(*lexmodelbuildingservice.GetUtterancesViewInput) lexmodelbuildingservice.GetUtterancesViewRequest

	PutBotRequest(*lexmodelbuildingservice.PutBotInput) lexmodelbuildingservice.PutBotRequest

	PutBotAliasRequest(*lexmodelbuildingservice.PutBotAliasInput) lexmodelbuildingservice.PutBotAliasRequest

	PutIntentRequest(*lexmodelbuildingservice.PutIntentInput) lexmodelbuildingservice.PutIntentRequest

	PutSlotTypeRequest(*lexmodelbuildingservice.PutSlotTypeInput) lexmodelbuildingservice.PutSlotTypeRequest
}

var _ LexModelBuildingServiceAPI = (*lexmodelbuildingservice.LexModelBuildingService)(nil)