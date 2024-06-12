// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: node_stake_update.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Updates the staking info at the end of staking period to indicate new staking period has started.
type NodeStakeUpdateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Time and date of the end of the staking period that is ending
	EndOfStakingPeriod *Timestamp `protobuf:"bytes,1,opt,name=end_of_staking_period,json=endOfStakingPeriod,proto3" json:"end_of_staking_period,omitempty"`
	// *
	// Staking info of each node at the beginning of the new staking period
	NodeStake []*NodeStake `protobuf:"bytes,2,rep,name=node_stake,json=nodeStake,proto3" json:"node_stake,omitempty"`
	// *
	// The maximum reward rate, in tinybars per whole hbar, that any account could receive in this
	// staking period.
	MaxStakingRewardRatePerHbar int64 `protobuf:"varint,3,opt,name=max_staking_reward_rate_per_hbar,json=maxStakingRewardRatePerHbar,proto3" json:"max_staking_reward_rate_per_hbar,omitempty"`
	// *
	// The fraction of the network and service fees paid to the node reward account 0.0.801.
	NodeRewardFeeFraction *Fraction `protobuf:"bytes,4,opt,name=node_reward_fee_fraction,json=nodeRewardFeeFraction,proto3" json:"node_reward_fee_fraction,omitempty"`
	// *
	// The maximum number of trailing periods for which a user can collect rewards. For example, if this
	// is 365 with a UTC calendar day period, then users must collect rewards at least once per calendar
	// year to avoid missing any value.
	StakingPeriodsStored int64 `protobuf:"varint,5,opt,name=staking_periods_stored,json=stakingPeriodsStored,proto3" json:"staking_periods_stored,omitempty"`
	// *
	// The number of minutes in a staking period. Note for the special case of 1440 minutes, periods are
	// treated as UTC calendar days, rather than repeating 1440 minute periods left-aligned at the epoch.
	StakingPeriod int64 `protobuf:"varint,6,opt,name=staking_period,json=stakingPeriod,proto3" json:"staking_period,omitempty"`
	// *
	// The fraction of the network and service fees paid to the staking reward account 0.0.800.
	StakingRewardFeeFraction *Fraction `protobuf:"bytes,7,opt,name=staking_reward_fee_fraction,json=stakingRewardFeeFraction,proto3" json:"staking_reward_fee_fraction,omitempty"`
	// *
	// The minimum balance of staking reward account 0.0.800 required to active rewards.
	StakingStartThreshold int64 `protobuf:"varint,8,opt,name=staking_start_threshold,json=stakingStartThreshold,proto3" json:"staking_start_threshold,omitempty"`
	// *
	// (DEPRECATED) The maximum total number of tinybars to be distributed as staking rewards in the
	// ending period. Please consult the max_total_reward field instead.
	//
	// Deprecated: Do not use.
	StakingRewardRate int64 `protobuf:"varint,9,opt,name=staking_reward_rate,json=stakingRewardRate,proto3" json:"staking_reward_rate,omitempty"`
	// *
	// The amount of the staking reward funds (account 0.0.800) reserved to pay pending rewards that
	// have been earned but not collected.
	ReservedStakingRewards int64 `protobuf:"varint,10,opt,name=reserved_staking_rewards,json=reservedStakingRewards,proto3" json:"reserved_staking_rewards,omitempty"`
	// *
	// The unreserved balance of account 0.0.800 at the close of the just-ending period; this value is
	// used to compute the HIP-782 balance ratio.
	UnreservedStakingRewardBalance int64 `protobuf:"varint,11,opt,name=unreserved_staking_reward_balance,json=unreservedStakingRewardBalance,proto3" json:"unreserved_staking_reward_balance,omitempty"`
	// *
	// The unreserved tinybar balance of account 0.0.800 required to achieve the maximum per-hbar reward
	// rate in any period; please see HIP-782 for details.
	RewardBalanceThreshold int64 `protobuf:"varint,12,opt,name=reward_balance_threshold,json=rewardBalanceThreshold,proto3" json:"reward_balance_threshold,omitempty"`
	// *
	// The maximum amount of tinybar that can be staked for reward while still achieving the maximum
	// per-hbar reward rate in any period; please see HIP-782 for details.
	MaxStakeRewarded int64 `protobuf:"varint,13,opt,name=max_stake_rewarded,json=maxStakeRewarded,proto3" json:"max_stake_rewarded,omitempty"`
	// *
	// The maximum total tinybars that could be paid as staking rewards in the ending period, after
	// applying the settings for the 0.0.800 balance threshold and the maximum stake rewarded. This
	// field replaces the deprecated field staking_reward_rate. It is only for convenience, since a
	// mirror node could also calculate its value by iterating the node_stake list and summing
	// stake_rewarded fields; then multiplying this sum by the max_staking_reward_rate_per_hbar.
	MaxTotalReward int64 `protobuf:"varint,14,opt,name=max_total_reward,json=maxTotalReward,proto3" json:"max_total_reward,omitempty"`
}

func (x *NodeStakeUpdateTransactionBody) Reset() {
	*x = NodeStakeUpdateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_stake_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStakeUpdateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStakeUpdateTransactionBody) ProtoMessage() {}

func (x *NodeStakeUpdateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_node_stake_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStakeUpdateTransactionBody.ProtoReflect.Descriptor instead.
func (*NodeStakeUpdateTransactionBody) Descriptor() ([]byte, []int) {
	return file_node_stake_update_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStakeUpdateTransactionBody) GetEndOfStakingPeriod() *Timestamp {
	if x != nil {
		return x.EndOfStakingPeriod
	}
	return nil
}

func (x *NodeStakeUpdateTransactionBody) GetNodeStake() []*NodeStake {
	if x != nil {
		return x.NodeStake
	}
	return nil
}

func (x *NodeStakeUpdateTransactionBody) GetMaxStakingRewardRatePerHbar() int64 {
	if x != nil {
		return x.MaxStakingRewardRatePerHbar
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetNodeRewardFeeFraction() *Fraction {
	if x != nil {
		return x.NodeRewardFeeFraction
	}
	return nil
}

func (x *NodeStakeUpdateTransactionBody) GetStakingPeriodsStored() int64 {
	if x != nil {
		return x.StakingPeriodsStored
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetStakingPeriod() int64 {
	if x != nil {
		return x.StakingPeriod
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetStakingRewardFeeFraction() *Fraction {
	if x != nil {
		return x.StakingRewardFeeFraction
	}
	return nil
}

func (x *NodeStakeUpdateTransactionBody) GetStakingStartThreshold() int64 {
	if x != nil {
		return x.StakingStartThreshold
	}
	return 0
}

// Deprecated: Do not use.
func (x *NodeStakeUpdateTransactionBody) GetStakingRewardRate() int64 {
	if x != nil {
		return x.StakingRewardRate
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetReservedStakingRewards() int64 {
	if x != nil {
		return x.ReservedStakingRewards
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetUnreservedStakingRewardBalance() int64 {
	if x != nil {
		return x.UnreservedStakingRewardBalance
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetRewardBalanceThreshold() int64 {
	if x != nil {
		return x.RewardBalanceThreshold
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetMaxStakeRewarded() int64 {
	if x != nil {
		return x.MaxStakeRewarded
	}
	return 0
}

func (x *NodeStakeUpdateTransactionBody) GetMaxTotalReward() int64 {
	if x != nil {
		return x.MaxTotalReward
	}
	return 0
}

// *
// Staking info for each node at the end of a staking period.
type NodeStake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The maximum stake (rewarded or not rewarded) this node can have as consensus weight. If its stake to
	// reward is above this maximum at the start of a period, then accounts staking to the node in that
	// period will be rewarded at a lower rate scaled by (maxStake / stakeRewardStart).
	MaxStake int64 `protobuf:"varint,1,opt,name=max_stake,json=maxStake,proto3" json:"max_stake,omitempty"`
	// *
	// The minimum stake (rewarded or not rewarded) this node must reach before having non-zero consensus weight.
	// If its total stake is below this minimum at the start of a period, then accounts staking to the node in
	// that period will receive no rewards.
	MinStake int64 `protobuf:"varint,2,opt,name=min_stake,json=minStake,proto3" json:"min_stake,omitempty"`
	// *
	// The id of this node.
	NodeId int64 `protobuf:"varint,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// *
	// The reward rate _per whole hbar_ that was staked to this node with declineReward=false from the start of
	// the staking period that is ending.
	RewardRate int64 `protobuf:"varint,4,opt,name=reward_rate,json=rewardRate,proto3" json:"reward_rate,omitempty"`
	// *
	// Consensus weight of this node for the new staking period.
	Stake int64 `protobuf:"varint,5,opt,name=stake,proto3" json:"stake,omitempty"`
	// *
	// Total of (balance + stakedToMe) for all accounts staked to this node with declineReward=true, at the
	// beginning of the new staking period.
	StakeNotRewarded int64 `protobuf:"varint,6,opt,name=stake_not_rewarded,json=stakeNotRewarded,proto3" json:"stake_not_rewarded,omitempty"`
	// *
	// Total of (balance + stakedToMe) for all accounts staked to this node with declineReward=false, at the
	// beginning of the new staking period.
	StakeRewarded int64 `protobuf:"varint,7,opt,name=stake_rewarded,json=stakeRewarded,proto3" json:"stake_rewarded,omitempty"`
}

func (x *NodeStake) Reset() {
	*x = NodeStake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_stake_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStake) ProtoMessage() {}

func (x *NodeStake) ProtoReflect() protoreflect.Message {
	mi := &file_node_stake_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStake.ProtoReflect.Descriptor instead.
func (*NodeStake) Descriptor() ([]byte, []int) {
	return file_node_stake_update_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStake) GetMaxStake() int64 {
	if x != nil {
		return x.MaxStake
	}
	return 0
}

func (x *NodeStake) GetMinStake() int64 {
	if x != nil {
		return x.MinStake
	}
	return 0
}

func (x *NodeStake) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeStake) GetRewardRate() int64 {
	if x != nil {
		return x.RewardRate
	}
	return 0
}

func (x *NodeStake) GetStake() int64 {
	if x != nil {
		return x.Stake
	}
	return 0
}

func (x *NodeStake) GetStakeNotRewarded() int64 {
	if x != nil {
		return x.StakeNotRewarded
	}
	return 0
}

func (x *NodeStake) GetStakeRewarded() int64 {
	if x != nil {
		return x.StakeRewarded
	}
	return 0
}

var File_node_stake_update_proto protoreflect.FileDescriptor

var file_node_stake_update_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x06, 0x0a, 0x1e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x43, 0x0a, 0x15, 0x65, 0x6e, 0x64, 0x5f, 0x6f,
	0x66, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x53,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2f, 0x0a, 0x0a,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x45, 0x0a,
	0x20, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x62, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x48, 0x62, 0x61, 0x72, 0x12, 0x48, 0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x46, 0x65, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x4e, 0x0a, 0x1b, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x65,
	0x65, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x18, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x46, 0x65, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x11, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x49, 0x0a, 0x21, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1e, 0x75, 0x6e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x18,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x6d, 0x61, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0xea,
	0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69,
	0x6e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x42, 0x26, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76,
	0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_stake_update_proto_rawDescOnce sync.Once
	file_node_stake_update_proto_rawDescData = file_node_stake_update_proto_rawDesc
)

func file_node_stake_update_proto_rawDescGZIP() []byte {
	file_node_stake_update_proto_rawDescOnce.Do(func() {
		file_node_stake_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_stake_update_proto_rawDescData)
	})
	return file_node_stake_update_proto_rawDescData
}

var file_node_stake_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_stake_update_proto_goTypes = []interface{}{
	(*NodeStakeUpdateTransactionBody)(nil), // 0: proto.NodeStakeUpdateTransactionBody
	(*NodeStake)(nil),                      // 1: proto.NodeStake
	(*Timestamp)(nil),                      // 2: proto.Timestamp
	(*Fraction)(nil),                       // 3: proto.Fraction
}
var file_node_stake_update_proto_depIdxs = []int32{
	2, // 0: proto.NodeStakeUpdateTransactionBody.end_of_staking_period:type_name -> proto.Timestamp
	1, // 1: proto.NodeStakeUpdateTransactionBody.node_stake:type_name -> proto.NodeStake
	3, // 2: proto.NodeStakeUpdateTransactionBody.node_reward_fee_fraction:type_name -> proto.Fraction
	3, // 3: proto.NodeStakeUpdateTransactionBody.staking_reward_fee_fraction:type_name -> proto.Fraction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_node_stake_update_proto_init() }
func file_node_stake_update_proto_init() {
	if File_node_stake_update_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_node_stake_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStakeUpdateTransactionBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_stake_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStake); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_stake_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_stake_update_proto_goTypes,
		DependencyIndexes: file_node_stake_update_proto_depIdxs,
		MessageInfos:      file_node_stake_update_proto_msgTypes,
	}.Build()
	File_node_stake_update_proto = out.File
	file_node_stake_update_proto_rawDesc = nil
	file_node_stake_update_proto_goTypes = nil
	file_node_stake_update_proto_depIdxs = nil
}
