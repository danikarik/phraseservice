package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

type createPhraseRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"creator"`
	Text    string       `json:"text"`
	Block   int64        `json:"block"`
}

func createPhraseHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPhraseRequest
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgCreatePhrase(owner, req.Text, req.Block)
		if err = msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, ctx, baseReq, []sdk.Msg{msg})
	}
}

type setPhraseRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"creator"`
	Text    string       `json:"text"`
	Block   int64        `json:"block"`
}

func setPhraseHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setPhraseRequest
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSetPhrase(owner, req.Text, req.Block)
		if err = msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, ctx, baseReq, []sdk.Msg{msg})
	}
}

type deletePhraseRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"creator"`
	Text    string       `json:"text"`
}

func deletePhraseHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deletePhraseRequest
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeletePhrase(owner, req.Text)
		if err = msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, ctx, baseReq, []sdk.Msg{msg})
	}
}
