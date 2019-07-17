// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleql

import (
	"time"

	"github.com/graphql-go/graphql"

	"storj.io/storj/internal/currency"
	"storj.io/storj/satellite/rewards"
)

const (
	// RewardType is a graphql type for reward
	RewardType = "reward"
	// RedeemRewardType is a graphql type for reward used for redemption of credits
	RedeemRewardType = "redeemReward"
	// FieldAwardCreditInCent is a field name for award credit amount for referrers
	FieldAwardCreditInCent = "awardCreditInCent"
	// FieldInviteeCreditInCents is a field name for credit amount rewarded to invitees
	FieldInviteeCreditInCents = "inviteeCreditInCents"
	// FieldRedeemableCap is a field name for the total redeemable amount of the reward offer
	FieldRedeemableCap = "redeemableCap"
	// FieldAwardCreditDurationDays is a field name for the valid time frame of current award credit
	FieldAwardCreditDurationDays = "awardCreditDurationDays"
	// FieldInviteeCreditDurationDays is a field name for the valid time frame of current invitee credit
	FieldInviteeCreditDurationDays = "inviteeCreditDurationDays"
	// FieldExpiresAt is a field name for the expiration time of a reward offer
	FieldExpiresAt = "expiresAt"
	// FieldType is a field name for the type of reward offers
	FieldType = "type"
	// FieldStatus is a field name for the status of reward offers
	FieldStatus = "status"
)

func graphqlReward() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: RewardType,
		Fields: graphql.Fields{
			FieldID: &graphql.Field{
				Type: graphql.Int,
			},
			FieldAwardCreditInCent: &graphql.Field{
				Type: graphql.Int,
			},
			FieldInviteeCreditInCents: &graphql.Field{
				Type: graphql.Int,
			},
			FieldRedeemableCap: &graphql.Field{
				Type: graphql.Int,
			},
			FieldAwardCreditDurationDays: &graphql.Field{
				Type: graphql.Int,
			},
			FieldInviteeCreditDurationDays: &graphql.Field{
				Type: graphql.Int,
			},
			FieldType: &graphql.Field{
				Type: graphql.Int,
			},
			FieldStatus: &graphql.Field{
				Type: graphql.Int,
			},
			FieldExpiresAt: &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func fromMapRewardInfo(args map[string]interface{}) (reward rewards.OfferInfo, err error) {
	reward.ID = args[FieldID].(int)
	if args[FieldInviteeCreditInCents] != nil {
		reward.InviteeCredit = currency.Cents(args[FieldInviteeCreditInCents].(int))
	}
	if args[FieldInviteeCreditDurationDays] != nil {
		reward.InviteeCreditDurationDays = args[FieldInviteeCreditDurationDays].(int)
	}
	if args[FieldAwardCreditInCent] != nil {
		reward.AwardCredit = currency.Cents(args[FieldAwardCreditInCent].(int))
	}
	if args[FieldAwardCreditDurationDays] != nil {
		reward.AwardCreditDurationDays = args[FieldAwardCreditDurationDays].(int)
	}
	reward.RedeemableCap = args[FieldRedeemableCap].(int)
	if args[FieldType] != nil {
		reward.Type = args[FieldType].(rewards.OfferType)
	}
	if args[FieldStatus] != nil {
		reward.Status = args[FieldStatus].(rewards.OfferStatus)
	}
	expiresAt, err := time.Parse(time.RFC3339, args[FieldExpiresAt].(string))
	if err != nil {
		return reward, err
	}
	reward.ExpiresAt = expiresAt
	return
}

func graphqlRedeemReward() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: RedeemRewardType,
		Fields: graphql.InputObjectConfigFieldMap{
			FieldID: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldAwardCreditInCent: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldInviteeCreditInCents: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldRedeemableCap: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldAwardCreditDurationDays: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldInviteeCreditDurationDays: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldType: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldStatus: &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			FieldExpiresAt: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
