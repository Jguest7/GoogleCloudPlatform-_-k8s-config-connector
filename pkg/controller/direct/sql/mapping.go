// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"fmt"

	api "google.golang.org/api/sqladmin/v1beta4"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
)

func SQLInstanceKRMToGCP(in *krm.SQLInstance, refs *SQLInstanceInternalRefs) (*api.DatabaseInstance, error) {
	out := &api.DatabaseInstance{}

	if in == nil {
		return nil, fmt.Errorf("cannot convert nil SQLInstance")
	}

	if in.Spec.DatabaseVersion != nil {
		out.DatabaseVersion = *in.Spec.DatabaseVersion
	}

	if in.Spec.EncryptionKMSCryptoKeyRef != nil {
		out.DiskEncryptionConfiguration = &api.DiskEncryptionConfiguration{
			Kind:       "sql#diskEncryptionConfiguration",
			KmsKeyName: refs.cryptoKey,
		}
	}

	if in.Spec.InstanceType != nil {
		out.InstanceType = *in.Spec.InstanceType
	}

	if in.Spec.MaintenanceVersion != nil {
		out.MaintenanceVersion = *in.Spec.MaintenanceVersion
	}

	if in.Spec.MasterInstanceRef != nil {
		out.MasterInstanceName = refs.masterInstance
	}

	if in.Spec.Region != nil {
		out.Region = *in.Spec.Region
	}

	if in.Spec.ReplicaConfiguration != nil {
		replicaConfiguration := &api.ReplicaConfiguration{
			Kind: "sql#replicaConfiguration",
		}

		if in.Spec.ReplicaConfiguration.FailoverTarget != nil {
			replicaConfiguration.FailoverTarget = *in.Spec.ReplicaConfiguration.FailoverTarget
		}

		// todo: requires mysql
		if in.Spec.ReplicaConfiguration.CaCertificate != nil ||
			in.Spec.ReplicaConfiguration.ClientCertificate != nil ||
			in.Spec.ReplicaConfiguration.ClientKey != nil ||
			in.Spec.ReplicaConfiguration.ConnectRetryInterval != nil ||
			in.Spec.ReplicaConfiguration.DumpFilePath != nil ||
			in.Spec.ReplicaConfiguration.MasterHeartbeatPeriod != nil ||
			in.Spec.ReplicaConfiguration.Password != nil ||
			in.Spec.ReplicaConfiguration.SslCipher != nil ||
			in.Spec.ReplicaConfiguration.Username != nil ||
			in.Spec.ReplicaConfiguration.VerifyServerCertificate != nil {
			replicaConfiguration.MysqlReplicaConfiguration = &api.MySqlReplicaConfiguration{}
		}

		if in.Spec.ReplicaConfiguration.CaCertificate != nil {
			replicaConfiguration.MysqlReplicaConfiguration.CaCertificate = *in.Spec.ReplicaConfiguration.CaCertificate
		}

		if in.Spec.ReplicaConfiguration.ClientCertificate != nil {
			replicaConfiguration.MysqlReplicaConfiguration.ClientCertificate = *in.Spec.ReplicaConfiguration.ClientCertificate
		}

		if in.Spec.ReplicaConfiguration.ClientKey != nil {
			replicaConfiguration.MysqlReplicaConfiguration.ClientKey = *in.Spec.ReplicaConfiguration.ClientKey
		}

		if in.Spec.ReplicaConfiguration.ConnectRetryInterval != nil {
			replicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval = *in.Spec.ReplicaConfiguration.ConnectRetryInterval
		}

		if in.Spec.ReplicaConfiguration.DumpFilePath != nil {
			replicaConfiguration.MysqlReplicaConfiguration.DumpFilePath = *in.Spec.ReplicaConfiguration.DumpFilePath
		}

		if in.Spec.ReplicaConfiguration.MasterHeartbeatPeriod != nil {
			replicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod = *in.Spec.ReplicaConfiguration.MasterHeartbeatPeriod
		}

		if in.Spec.ReplicaConfiguration.Password != nil {
			replicaConfiguration.MysqlReplicaConfiguration.Password = refs.replicaPassword
		}

		if in.Spec.ReplicaConfiguration.SslCipher != nil {
			replicaConfiguration.MysqlReplicaConfiguration.SslCipher = *in.Spec.ReplicaConfiguration.SslCipher
		}

		if in.Spec.ReplicaConfiguration.Username != nil {
			replicaConfiguration.MysqlReplicaConfiguration.Username = *in.Spec.ReplicaConfiguration.Username
		}

		if in.Spec.ReplicaConfiguration.VerifyServerCertificate != nil {
			replicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate = *in.Spec.ReplicaConfiguration.VerifyServerCertificate
		}

		out.ReplicaConfiguration = replicaConfiguration
	}

	if in.Spec.ResourceID != nil {
		out.Name = *in.Spec.ResourceID
	} else {
		return nil, fmt.Errorf("resourceID is empty")
	}

	if in.Spec.RootPassword != nil && refs.rootPassword != "" {
		out.RootPassword = refs.rootPassword
	}

	out.Settings = &api.Settings{}

	if in.Spec.Settings.ActivationPolicy != nil {
		out.Settings.ActivationPolicy = *in.Spec.Settings.ActivationPolicy
	}

	if in.Spec.Settings.ActiveDirectoryConfig != nil {
		// todo: requires sqlserver
		// todo: requires private network
		out.Settings.ActiveDirectoryConfig = &api.SqlActiveDirectoryConfig{
			Domain: in.Spec.Settings.ActiveDirectoryConfig.Domain,
		}
	}

	if in.Spec.Settings.AdvancedMachineFeatures != nil {
		// todo: requires sqlserver
		// todo: requires >= 6 cpu cores
		out.Settings.AdvancedMachineFeatures = &api.AdvancedMachineFeatures{}
		if in.Spec.Settings.AdvancedMachineFeatures.ThreadsPerCore != nil {
			out.Settings.AdvancedMachineFeatures.ThreadsPerCore = *in.Spec.Settings.AdvancedMachineFeatures.ThreadsPerCore
		}
	}

	if in.Spec.Settings.AuthorizedGaeApplications != nil {
		// todo: deprecated
		out.Settings.AuthorizedGaeApplications = in.Spec.Settings.AuthorizedGaeApplications
	}

	if in.Spec.Settings.AvailabilityType != nil {
		out.Settings.AvailabilityType = *in.Spec.Settings.AvailabilityType
	}

	if in.Spec.Settings.BackupConfiguration != nil {
		backupConfig := &api.BackupConfiguration{}

		if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings != nil {
			retentionSettings := &api.BackupRetentionSettings{
				RetainedBackups: in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups,
			}
			if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit != nil {
				retentionSettings.RetentionUnit = *in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit
			}
			backupConfig.BackupRetentionSettings = retentionSettings
		}

		if in.Spec.Settings.BackupConfiguration.BinaryLogEnabled != nil {
			// todo: requires mysql
			backupConfig.BinaryLogEnabled = *in.Spec.Settings.BackupConfiguration.BinaryLogEnabled
		}

		if in.Spec.Settings.BackupConfiguration.Enabled != nil {
			backupConfig.Enabled = *in.Spec.Settings.BackupConfiguration.Enabled
		}

		if in.Spec.Settings.BackupConfiguration.Location != nil {
			backupConfig.Location = *in.Spec.Settings.BackupConfiguration.Location
		}

		if in.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled != nil {
			backupConfig.PointInTimeRecoveryEnabled = *in.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled
		}

		if in.Spec.Settings.BackupConfiguration.StartTime != nil {
			backupConfig.StartTime = *in.Spec.Settings.BackupConfiguration.StartTime
		}

		if in.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays != nil {
			backupConfig.TransactionLogRetentionDays = *in.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays
		}

		out.Settings.BackupConfiguration = backupConfig
	}

	if in.Spec.Settings.Collation != nil {
		// todo: requires sqlserver
		out.Settings.Collation = *in.Spec.Settings.Collation
	}

	if in.Spec.Settings.ConnectorEnforcement != nil {
		out.Settings.ConnectorEnforcement = *in.Spec.Settings.ConnectorEnforcement
	}

	if in.Spec.Settings.CrashSafeReplication != nil {
		// todo: deprecated
		out.Settings.CrashSafeReplicationEnabled = *in.Spec.Settings.CrashSafeReplication
	}

	if in.Spec.Settings.DataCacheConfig != nil {
		// todo: requires ENTERPRISE_PLUS edition
		// todo: requires mysql or postgres
		out.Settings.DataCacheConfig = &api.DataCacheConfig{
			DataCacheEnabled: *in.Spec.Settings.DataCacheConfig.DataCacheEnabled,
		}
	}

	if in.Spec.Settings.DatabaseFlags != nil {
		dbFlags := []*api.DatabaseFlags{}
		for _, dbFlag := range in.Spec.Settings.DatabaseFlags {
			dbFlags = append(dbFlags, &api.DatabaseFlags{
				Name:  dbFlag.Name,
				Value: dbFlag.Value,
			})
		}
		out.Settings.DatabaseFlags = dbFlags
	}

	if in.Spec.Settings.DeletionProtectionEnabled != nil {
		out.Settings.DeletionProtectionEnabled = *in.Spec.Settings.DeletionProtectionEnabled
	}

	if in.Spec.Settings.DenyMaintenancePeriod != nil {
		// todo: handle multiple periods (?)
		out.Settings.DenyMaintenancePeriods = []*api.DenyMaintenancePeriod{
			{
				EndDate:   in.Spec.Settings.DenyMaintenancePeriod.EndDate,
				StartDate: in.Spec.Settings.DenyMaintenancePeriod.StartDate,
				Time:      in.Spec.Settings.DenyMaintenancePeriod.Time,
			},
		}
	}

	if in.Spec.Settings.DiskAutoresize != nil {
		out.Settings.StorageAutoResize = in.Spec.Settings.DiskAutoresize
	}

	if in.Spec.Settings.DiskAutoresizeLimit != nil {
		// todo: requires DiskAutoresize == true
		out.Settings.StorageAutoResizeLimit = *in.Spec.Settings.DiskAutoresizeLimit
	}

	if in.Spec.Settings.DiskSize != nil {
		out.Settings.DataDiskSizeGb = *in.Spec.Settings.DiskSize
	}

	if in.Spec.Settings.DiskType != nil {
		out.Settings.DataDiskType = *in.Spec.Settings.DiskType
	}

	if in.Spec.Settings.Edition != nil {
		out.Settings.Edition = *in.Spec.Settings.Edition
	}

	if in.Spec.Settings.InsightsConfig != nil {
		insightsConfig := &api.InsightsConfig{}

		if in.Spec.Settings.InsightsConfig.QueryInsightsEnabled != nil {
			insightsConfig.QueryInsightsEnabled = *in.Spec.Settings.InsightsConfig.QueryInsightsEnabled
		}

		if in.Spec.Settings.InsightsConfig.QueryPlansPerMinute != nil {
			insightsConfig.QueryPlansPerMinute = *in.Spec.Settings.InsightsConfig.QueryPlansPerMinute
		}

		if in.Spec.Settings.InsightsConfig.QueryStringLength != nil {
			insightsConfig.QueryStringLength = *in.Spec.Settings.InsightsConfig.QueryStringLength
		}

		if in.Spec.Settings.InsightsConfig.RecordApplicationTags != nil {
			insightsConfig.RecordApplicationTags = *in.Spec.Settings.InsightsConfig.RecordApplicationTags
		}

		if in.Spec.Settings.InsightsConfig.RecordClientAddress != nil {
			insightsConfig.RecordClientAddress = *in.Spec.Settings.InsightsConfig.RecordClientAddress
		}

		out.Settings.InsightsConfig = insightsConfig
	}

	if in.Spec.Settings.IpConfiguration != nil {
		out.Settings.IpConfiguration = &api.IpConfiguration{}

		if in.Spec.Settings.IpConfiguration.AllocatedIpRange != nil {
			out.Settings.IpConfiguration.AllocatedIpRange = *in.Spec.Settings.IpConfiguration.AllocatedIpRange
		}

		if in.Spec.Settings.IpConfiguration.AuthorizedNetworks != nil {
			authorizedNetworks := []*api.AclEntry{}
			for _, net := range in.Spec.Settings.IpConfiguration.AuthorizedNetworks {
				netEntry := &api.AclEntry{
					Kind:  "sql#aclEntry",
					Value: net.Value,
				}
				if net.ExpirationTime != nil {
					netEntry.ExpirationTime = *net.ExpirationTime
				}
				if net.Name != nil {
					netEntry.Name = *net.Name
				}
				authorizedNetworks = append(authorizedNetworks, netEntry)
			}
			out.Settings.IpConfiguration.AuthorizedNetworks = authorizedNetworks
		}

		if in.Spec.Settings.IpConfiguration.EnablePrivatePathForGoogleCloudServices != nil {
			out.Settings.IpConfiguration.EnablePrivatePathForGoogleCloudServices = *in.Spec.Settings.IpConfiguration.EnablePrivatePathForGoogleCloudServices
		}

		if in.Spec.Settings.IpConfiguration.Ipv4Enabled != nil {
			out.Settings.IpConfiguration.Ipv4Enabled = *in.Spec.Settings.IpConfiguration.Ipv4Enabled
		} else {
			out.Settings.IpConfiguration.Ipv4Enabled = false
		}
		// todo: must do this for all fields (?)
		out.Settings.IpConfiguration.ForceSendFields = []string{"Ipv4Enabled"}

		if in.Spec.Settings.IpConfiguration.PrivateNetworkRef != nil {
			out.Settings.IpConfiguration.PrivateNetwork = refs.privateNetwork
		}

		if in.Spec.Settings.IpConfiguration.PscConfig != nil {
			if len(in.Spec.Settings.IpConfiguration.PscConfig) != 1 {
				return nil, fmt.Errorf("only one psc config allowed per instance")
			}
			out.Settings.IpConfiguration.PscConfig = &api.PscConfig{
				AllowedConsumerProjects: in.Spec.Settings.IpConfiguration.PscConfig[0].AllowedConsumerProjects,
			}
			if in.Spec.Settings.IpConfiguration.PscConfig[0].PscEnabled != nil {
				out.Settings.IpConfiguration.PscConfig.PscEnabled = *in.Spec.Settings.IpConfiguration.PscConfig[0].PscEnabled
			}
		}

		if in.Spec.Settings.IpConfiguration.RequireSsl != nil {
			// todo: deprecated
			out.Settings.IpConfiguration.RequireSsl = *in.Spec.Settings.IpConfiguration.RequireSsl
		}

		if in.Spec.Settings.IpConfiguration.SslMode != nil {
			out.Settings.IpConfiguration.SslMode = *in.Spec.Settings.IpConfiguration.SslMode
		}
	}

	if in.Spec.Settings.LocationPreference != nil {
		out.Settings.LocationPreference = &api.LocationPreference{}
		if in.Spec.Settings.LocationPreference.FollowGaeApplication != nil {
			// todo: deprecated
			out.Settings.LocationPreference.FollowGaeApplication = *in.Spec.Settings.LocationPreference.FollowGaeApplication
		}
		if in.Spec.Settings.LocationPreference.SecondaryZone != nil {
			out.Settings.LocationPreference.SecondaryZone = *in.Spec.Settings.LocationPreference.SecondaryZone
		}
		if in.Spec.Settings.LocationPreference.Zone != nil {
			out.Settings.LocationPreference.Zone = *in.Spec.Settings.LocationPreference.Zone
		}
	}

	if in.Spec.Settings.MaintenanceWindow != nil {
		out.Settings.MaintenanceWindow = &api.MaintenanceWindow{}
		if in.Spec.Settings.MaintenanceWindow.Day != nil {
			out.Settings.MaintenanceWindow.Day = *in.Spec.Settings.MaintenanceWindow.Day
		}
		if in.Spec.Settings.MaintenanceWindow.Hour != nil {
			out.Settings.MaintenanceWindow.Hour = *in.Spec.Settings.MaintenanceWindow.Hour
		}
		if in.Spec.Settings.MaintenanceWindow.UpdateTrack != nil {
			out.Settings.MaintenanceWindow.UpdateTrack = *in.Spec.Settings.MaintenanceWindow.UpdateTrack
		}
	}

	if in.Spec.Settings.PasswordValidationPolicy != nil {
		policy := &api.PasswordValidationPolicy{
			EnablePasswordPolicy: in.Spec.Settings.PasswordValidationPolicy.EnablePasswordPolicy,
		}

		if in.Spec.Settings.PasswordValidationPolicy.Complexity != nil {
			policy.Complexity = *in.Spec.Settings.PasswordValidationPolicy.Complexity
		}

		if in.Spec.Settings.PasswordValidationPolicy.DisallowUsernameSubstring != nil {
			policy.DisallowUsernameSubstring = *in.Spec.Settings.PasswordValidationPolicy.DisallowUsernameSubstring
		}

		if in.Spec.Settings.PasswordValidationPolicy.MinLength != nil {
			policy.MinLength = *in.Spec.Settings.PasswordValidationPolicy.MinLength
		}

		if in.Spec.Settings.PasswordValidationPolicy.PasswordChangeInterval != nil {
			policy.PasswordChangeInterval = *in.Spec.Settings.PasswordValidationPolicy.PasswordChangeInterval
		}

		if in.Spec.Settings.PasswordValidationPolicy.ReuseInterval != nil {
			policy.ReuseInterval = *in.Spec.Settings.PasswordValidationPolicy.ReuseInterval
		}

		out.Settings.PasswordValidationPolicy = policy
	}

	if in.Spec.Settings.PricingPlan != nil {
		// todo: can only be PER_USE
		out.Settings.PricingPlan = *in.Spec.Settings.PricingPlan
	}

	if in.Spec.Settings.ReplicationType != nil {
		// todo: deprecated
		out.Settings.ReplicationType = *in.Spec.Settings.ReplicationType
	}

	if in.Spec.Settings.SqlServerAuditConfig != nil {
		// todo: requires sqlserver
		out.Settings.SqlServerAuditConfig = &api.SqlServerAuditConfig{
			Kind: "sql#sqlServerAuditConfig",
		}

		if in.Spec.Settings.SqlServerAuditConfig.BucketRef != nil {
			// todo: required
			out.Settings.SqlServerAuditConfig.Bucket = refs.auditLogBucket
		}

		if in.Spec.Settings.SqlServerAuditConfig.RetentionInterval != nil {
			out.Settings.SqlServerAuditConfig.RetentionInterval = *in.Spec.Settings.SqlServerAuditConfig.RetentionInterval
		}

		if in.Spec.Settings.SqlServerAuditConfig.UploadInterval != nil {
			out.Settings.SqlServerAuditConfig.UploadInterval = *in.Spec.Settings.SqlServerAuditConfig.UploadInterval
		}
	}

	out.Settings.Tier = in.Spec.Settings.Tier

	if in.Spec.Settings.TimeZone != nil {
		// todo: requires sqlserver
		out.Settings.TimeZone = *in.Spec.Settings.TimeZone
	}

	if in.Labels != nil {
		out.Settings.UserLabels = make(map[string]string)
		for k, v := range in.Labels {
			out.Settings.UserLabels[k] = v
		}
	}

	return out, nil
}

func SQLInstanceGCPToKRM(in *api.DatabaseInstance) (*krm.SQLInstance, error) {
	out := &krm.SQLInstance{}

	if in == nil {
		return nil, fmt.Errorf("cannot convert nil DatabaseInstance")
	}

	if in.DatabaseVersion != "" {
		out.Spec.DatabaseVersion = &in.DatabaseVersion
	}

	if in.DiskEncryptionConfiguration != nil {
		out.Spec.EncryptionKMSCryptoKeyRef = &refs.KMSCryptoKeyRef{
			External: in.DiskEncryptionConfiguration.KmsKeyName,
		}
	}

	if in.InstanceType != "" {
		out.Spec.InstanceType = &in.InstanceType
	}

	if in.MaintenanceVersion != "" {
		out.Spec.MaintenanceVersion = &in.MaintenanceVersion
	}

	if in.MasterInstanceName != "" {
		out.Spec.MasterInstanceRef = &refs.SQLInstanceRef{
			External: in.MasterInstanceName,
		}
	}

	if in.Region != "" {
		out.Spec.Region = &in.Region
	}

	if in.ReplicaConfiguration != nil {
		rc := &krm.InstanceReplicaConfiguration{}

		rc.FailoverTarget = &in.ReplicaConfiguration.FailoverTarget

		if in.ReplicaConfiguration.MysqlReplicaConfiguration != nil {
			rc.CaCertificate = &in.ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate
			rc.ClientKey = &in.ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey
			rc.ConnectRetryInterval = &in.ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval
			rc.DumpFilePath = &in.ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath
			rc.MasterHeartbeatPeriod = &in.ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod
			rc.Password = &krm.InstancePassword{
				Value: &in.ReplicaConfiguration.MysqlReplicaConfiguration.Password,
			}
			rc.SslCipher = &in.ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher
			rc.Username = &in.ReplicaConfiguration.MysqlReplicaConfiguration.Username
			rc.VerifyServerCertificate = &in.ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate
		}

		out.Spec.ReplicaConfiguration = rc
	}

	out.Spec.ResourceID = &in.Name

	out.Spec.RootPassword = &krm.InstanceRootPassword{
		Value: &in.RootPassword,
	}

	if in.Settings.ActivationPolicy != "" {
		out.Spec.Settings.ActivationPolicy = &in.Settings.ActivationPolicy
	}

	if in.Settings.ActiveDirectoryConfig != nil {
		out.Spec.Settings.ActiveDirectoryConfig = &krm.InstanceActiveDirectoryConfig{
			Domain: in.Settings.ActiveDirectoryConfig.Domain,
		}
	}

	if in.Settings.AdvancedMachineFeatures != nil {
		out.Spec.Settings.AdvancedMachineFeatures = &krm.InstanceAdvancedMachineFeatures{
			ThreadsPerCore: &in.Settings.AdvancedMachineFeatures.ThreadsPerCore,
		}
	}

	if in.Settings.AuthorizedGaeApplications != nil {
		out.Spec.Settings.AuthorizedGaeApplications = in.Settings.AuthorizedGaeApplications
	}

	if in.Settings.AvailabilityType != "" {
		out.Spec.Settings.AvailabilityType = &in.Settings.AvailabilityType
	}

	if in.Settings.BackupConfiguration != nil {
		bc := &krm.InstanceBackupConfiguration{}

		if in.Settings.BackupConfiguration.BackupRetentionSettings != nil {
			bc.BackupRetentionSettings = &krm.InstanceBackupRetentionSettings{
				RetainedBackups: in.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups,
				RetentionUnit:   &in.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit,
			}
		}

		bc.BinaryLogEnabled = &in.Settings.BackupConfiguration.BinaryLogEnabled
		bc.Enabled = &in.Settings.BackupConfiguration.Enabled
		bc.Location = &in.Settings.BackupConfiguration.Location
		bc.PointInTimeRecoveryEnabled = &in.Settings.BackupConfiguration.PointInTimeRecoveryEnabled
		bc.StartTime = &in.Settings.BackupConfiguration.StartTime
		bc.TransactionLogRetentionDays = &in.Settings.BackupConfiguration.TransactionLogRetentionDays

		out.Spec.Settings.BackupConfiguration = bc
	}

	if in.Settings.Collation != "" {
		out.Spec.Settings.Collation = &in.Settings.Collation
	}

	if in.Settings.ConnectorEnforcement != "" {
		out.Spec.Settings.ConnectorEnforcement = &in.Settings.ConnectorEnforcement
	}

	out.Spec.Settings.CrashSafeReplication = &in.Settings.CrashSafeReplicationEnabled

	if in.Settings.DatabaseFlags != nil {
		dbFlags := []krm.InstanceDatabaseFlags{}
		for _, dbFlag := range in.Settings.DatabaseFlags {
			dbFlags = append(dbFlags, krm.InstanceDatabaseFlags{
				Name:  dbFlag.Name,
				Value: dbFlag.Value,
			})
		}
		out.Spec.Settings.DatabaseFlags = dbFlags
	}

	out.Spec.Settings.DeletionProtectionEnabled = &in.Settings.DeletionProtectionEnabled

	// todo: handle multiple periods (?)
	if in.Settings.DenyMaintenancePeriods != nil && len(in.Settings.DenyMaintenancePeriods) >= 1 {
		out.Spec.Settings.DenyMaintenancePeriod = &krm.InstanceDenyMaintenancePeriod{
			EndDate:   in.Settings.DenyMaintenancePeriods[0].EndDate,
			StartDate: in.Settings.DenyMaintenancePeriods[0].StartDate,
			Time:      in.Settings.DenyMaintenancePeriods[0].Time,
		}
	}

	out.Spec.Settings.DiskAutoresize = in.Settings.StorageAutoResize

	if in.Settings.StorageAutoResizeLimit != 0 {
		out.Spec.Settings.DiskAutoresizeLimit = &in.Settings.StorageAutoResizeLimit
	}

	out.Spec.Settings.DiskSize = &in.Settings.DataDiskSizeGb

	if in.Settings.DataDiskType != "" {
		out.Spec.Settings.DiskType = &in.Settings.DataDiskType
	}

	if in.Settings.Edition != "" {
		out.Spec.Settings.Edition = &in.Settings.Edition
	}

	if in.Settings.InsightsConfig != nil {
		out.Spec.Settings.InsightsConfig = &krm.InstanceInsightsConfig{
			QueryInsightsEnabled:  &in.Settings.InsightsConfig.QueryInsightsEnabled,
			QueryPlansPerMinute:   &in.Settings.InsightsConfig.QueryPlansPerMinute,
			QueryStringLength:     &in.Settings.InsightsConfig.QueryStringLength,
			RecordApplicationTags: &in.Settings.InsightsConfig.RecordApplicationTags,
			RecordClientAddress:   &in.Settings.InsightsConfig.RecordClientAddress,
		}
	}

	if in.Settings.IpConfiguration != nil {
		ic := &krm.InstanceIpConfiguration{}

		if in.Settings.IpConfiguration.AllocatedIpRange != "" {
			ic.AllocatedIpRange = &in.Settings.IpConfiguration.AllocatedIpRange
		}

		if in.Settings.IpConfiguration.AuthorizedNetworks != nil {
			ans := []krm.InstanceAuthorizedNetworks{}
			for _, an := range in.Settings.IpConfiguration.AuthorizedNetworks {
				ans = append(ans, krm.InstanceAuthorizedNetworks{
					ExpirationTime: &an.ExpirationTime,
					Name:           &an.Name,
					Value:          an.Value,
				})
			}
			ic.AuthorizedNetworks = ans
		}

		ic.EnablePrivatePathForGoogleCloudServices = &in.Settings.IpConfiguration.EnablePrivatePathForGoogleCloudServices

		ic.Ipv4Enabled = &in.Settings.IpConfiguration.Ipv4Enabled

		if in.Settings.IpConfiguration.PrivateNetwork != "" {
			ic.PrivateNetworkRef = &refs.ComputeNetworkRef{
				External: in.Settings.IpConfiguration.PrivateNetwork,
			}
		}

		if in.Settings.IpConfiguration.PscConfig != nil {
			out.Spec.Settings.IpConfiguration.PscConfig = []krm.InstancePscConfig{
				{
					AllowedConsumerProjects: in.Settings.IpConfiguration.PscConfig.AllowedConsumerProjects,
					PscEnabled:              &in.Settings.IpConfiguration.PscConfig.PscEnabled,
				},
			}
		}

		ic.RequireSsl = &in.Settings.IpConfiguration.RequireSsl

		if in.Settings.IpConfiguration.SslMode != "" {
			ic.SslMode = &in.Settings.IpConfiguration.SslMode
		}

		out.Spec.Settings.IpConfiguration = ic
	}

	if in.Settings.LocationPreference != nil {
		out.Spec.Settings.LocationPreference = &krm.InstanceLocationPreference{
			FollowGaeApplication: &in.Settings.LocationPreference.FollowGaeApplication,
			SecondaryZone:        &in.Settings.LocationPreference.SecondaryZone,
			Zone:                 &in.Settings.LocationPreference.Zone,
		}
	}

	if in.Settings.MaintenanceWindow != nil {
		out.Spec.Settings.MaintenanceWindow = &krm.InstanceMaintenanceWindow{
			Day:         &in.Settings.MaintenanceWindow.Day,
			Hour:        &in.Settings.MaintenanceWindow.Hour,
			UpdateTrack: &in.Settings.MaintenanceWindow.UpdateTrack,
		}
	}

	if in.Settings.PasswordValidationPolicy != nil {
		out.Spec.Settings.PasswordValidationPolicy = &krm.InstancePasswordValidationPolicy{
			EnablePasswordPolicy:      in.Settings.PasswordValidationPolicy.EnablePasswordPolicy,
			Complexity:                &in.Settings.PasswordValidationPolicy.Complexity,
			DisallowUsernameSubstring: &in.Settings.PasswordValidationPolicy.DisallowUsernameSubstring,
			MinLength:                 &in.Settings.PasswordValidationPolicy.MinLength,
			PasswordChangeInterval:    &in.Settings.PasswordValidationPolicy.PasswordChangeInterval,
			ReuseInterval:             &in.Settings.PasswordValidationPolicy.ReuseInterval,
		}
	}

	if in.Settings.PricingPlan != "" {
		out.Spec.Settings.PricingPlan = &in.Settings.PricingPlan
	}

	if in.Settings.ReplicationType != "" {
		out.Spec.Settings.ReplicationType = &in.Settings.ReplicationType
	}

	if in.Settings.SqlServerAuditConfig != nil {
		out.Spec.Settings.SqlServerAuditConfig = &krm.InstanceSqlServerAuditConfig{
			BucketRef: &refs.StorageBucketRef{
				External: in.Settings.SqlServerAuditConfig.Bucket,
			},
			RetentionInterval: &in.Settings.SqlServerAuditConfig.RetentionInterval,
			UploadInterval:    &in.Settings.SqlServerAuditConfig.UploadInterval,
		}
	}

	out.Spec.Settings.Tier = in.Settings.Tier

	if in.Settings.TimeZone != "" {
		out.Spec.Settings.TimeZone = &in.Settings.TimeZone
	}

	if in.Settings.UserLabels != nil {
		out.Labels = in.Settings.UserLabels
	}

	return out, nil
}

func Convert_SQLInstance_API_v1_To_KRM_status(in *api.DatabaseInstance, out *krm.SQLInstanceStatus) error {
	if in == nil {
		return fmt.Errorf("cannot convert nil DatabaseInstance")
	}

	if in.AvailableMaintenanceVersions != nil {
		out.AvailableMaintenanceVersions = append(out.AvailableMaintenanceVersions, in.AvailableMaintenanceVersions...)
	}

	out.ConnectionName = LazyPtr(in.ConnectionName)

	out.DnsName = LazyPtr(in.DnsName)

	if len(in.IpAddresses) >= 1 {
		out.FirstIpAddress = LazyPtr(in.IpAddresses[0].IpAddress)
	}

	out.InstanceType = LazyPtr(in.InstanceType)

	if in.IpAddresses != nil {
		for _, ia := range in.IpAddresses {
			ipAddr := krm.InstanceIpAddressStatus{
				IpAddress:    LazyPtr(ia.IpAddress),
				TimeToRetire: LazyPtr(ia.TimeToRetire),
				Type:         LazyPtr(ia.Type),
			}
			out.IpAddress = append(out.IpAddress, ipAddr)

			if ia.Type == "PRIMARY" {
				out.PublicIpAddress = LazyPtr(ia.IpAddress)
			}

			if ia.Type == "PRIVATE" {
				out.PrivateIpAddress = LazyPtr(ia.IpAddress)
			}
		}
	}

	out.PscServiceAttachmentLink = LazyPtr(in.PscServiceAttachmentLink)

	out.SelfLink = LazyPtr(in.SelfLink)

	if in.ServerCaCert != nil {
		out.ServerCaCert = &krm.InstanceServerCaCertStatus{
			Cert:            LazyPtr(in.ServerCaCert.Cert),
			CommonName:      LazyPtr(in.ServerCaCert.CommonName),
			CreateTime:      LazyPtr(in.ServerCaCert.CreateTime),
			ExpirationTime:  LazyPtr(in.ServerCaCert.ExpirationTime),
			Sha1Fingerprint: LazyPtr(in.ServerCaCert.Sha1Fingerprint),
		}
	}

	out.ServiceAccountEmailAddress = LazyPtr(in.ServiceAccountEmailAddress)

	return nil
}

func LazyPtr[T comparable](v T) *T {
	var defaultValue T
	if v == defaultValue {
		return nil
	}
	return &v
}
