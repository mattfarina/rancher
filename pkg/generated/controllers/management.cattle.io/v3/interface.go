/*
Copyright 2022 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v3

import (
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v3.AddToScheme)
}

type Interface interface {
	ActiveDirectoryProvider() ActiveDirectoryProviderController
	AuthConfig() AuthConfigController
	AuthProvider() AuthProviderController
	AuthToken() AuthTokenController
	AzureADProvider() AzureADProviderController
	Catalog() CatalogController
	CatalogTemplate() CatalogTemplateController
	CatalogTemplateVersion() CatalogTemplateVersionController
	CisBenchmarkVersion() CisBenchmarkVersionController
	CisConfig() CisConfigController
	CloudCredential() CloudCredentialController
	Cluster() ClusterController
	ClusterAlert() ClusterAlertController
	ClusterAlertGroup() ClusterAlertGroupController
	ClusterAlertRule() ClusterAlertRuleController
	ClusterCatalog() ClusterCatalogController
	ClusterLogging() ClusterLoggingController
	ClusterMonitorGraph() ClusterMonitorGraphController
	ClusterRegistrationToken() ClusterRegistrationTokenController
	ClusterRoleTemplateBinding() ClusterRoleTemplateBindingController
	ClusterScan() ClusterScanController
	ClusterTemplate() ClusterTemplateController
	ClusterTemplateRevision() ClusterTemplateRevisionController
	ComposeConfig() ComposeConfigController
	DynamicSchema() DynamicSchemaController
	EtcdBackup() EtcdBackupController
	Feature() FeatureController
	FleetWorkspace() FleetWorkspaceController
	FreeIpaProvider() FreeIpaProviderController
	GithubProvider() GithubProviderController
	GlobalDns() GlobalDnsController
	GlobalDnsProvider() GlobalDnsProviderController
	GlobalRole() GlobalRoleController
	GlobalRoleBinding() GlobalRoleBindingController
	GoogleOAuthProvider() GoogleOAuthProviderController
	Group() GroupController
	GroupMember() GroupMemberController
	KontainerDriver() KontainerDriverController
	LocalProvider() LocalProviderController
	MonitorMetric() MonitorMetricController
	MultiClusterApp() MultiClusterAppController
	MultiClusterAppRevision() MultiClusterAppRevisionController
	Node() NodeController
	NodeDriver() NodeDriverController
	NodePool() NodePoolController
	NodeTemplate() NodeTemplateController
	Notifier() NotifierController
	OpenLdapProvider() OpenLdapProviderController
	PodSecurityPolicyTemplate() PodSecurityPolicyTemplateController
	PodSecurityPolicyTemplateProjectBinding() PodSecurityPolicyTemplateProjectBindingController
	Preference() PreferenceController
	Principal() PrincipalController
	Project() ProjectController
	ProjectAlert() ProjectAlertController
	ProjectAlertGroup() ProjectAlertGroupController
	ProjectAlertRule() ProjectAlertRuleController
	ProjectCatalog() ProjectCatalogController
	ProjectLogging() ProjectLoggingController
	ProjectMonitorGraph() ProjectMonitorGraphController
	ProjectNetworkPolicy() ProjectNetworkPolicyController
	ProjectRoleTemplateBinding() ProjectRoleTemplateBindingController
	RkeAddon() RkeAddonController
	RkeK8sServiceOption() RkeK8sServiceOptionController
	RkeK8sSystemImage() RkeK8sSystemImageController
	RoleTemplate() RoleTemplateController
	SamlProvider() SamlProviderController
	SamlToken() SamlTokenController
	Setting() SettingController
	Template() TemplateController
	TemplateContent() TemplateContentController
	TemplateVersion() TemplateVersionController
	Token() TokenController
	User() UserController
	UserAttribute() UserAttributeController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) ActiveDirectoryProvider() ActiveDirectoryProviderController {
	return NewActiveDirectoryProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ActiveDirectoryProvider"}, "activedirectoryproviders", false, c.controllerFactory)
}
func (c *version) AuthConfig() AuthConfigController {
	return NewAuthConfigController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthConfig"}, "authconfigs", false, c.controllerFactory)
}
func (c *version) AuthProvider() AuthProviderController {
	return NewAuthProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthProvider"}, "authproviders", false, c.controllerFactory)
}
func (c *version) AuthToken() AuthTokenController {
	return NewAuthTokenController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthToken"}, "authtokens", false, c.controllerFactory)
}
func (c *version) AzureADProvider() AzureADProviderController {
	return NewAzureADProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AzureADProvider"}, "azureadproviders", false, c.controllerFactory)
}
func (c *version) Catalog() CatalogController {
	return NewCatalogController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Catalog"}, "catalogs", false, c.controllerFactory)
}
func (c *version) CatalogTemplate() CatalogTemplateController {
	return NewCatalogTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CatalogTemplate"}, "catalogtemplates", true, c.controllerFactory)
}
func (c *version) CatalogTemplateVersion() CatalogTemplateVersionController {
	return NewCatalogTemplateVersionController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CatalogTemplateVersion"}, "catalogtemplateversions", true, c.controllerFactory)
}
func (c *version) CisBenchmarkVersion() CisBenchmarkVersionController {
	return NewCisBenchmarkVersionController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CisBenchmarkVersion"}, "cisbenchmarkversions", true, c.controllerFactory)
}
func (c *version) CisConfig() CisConfigController {
	return NewCisConfigController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CisConfig"}, "cisconfigs", true, c.controllerFactory)
}
func (c *version) CloudCredential() CloudCredentialController {
	return NewCloudCredentialController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "CloudCredential"}, "cloudcredentials", true, c.controllerFactory)
}
func (c *version) Cluster() ClusterController {
	return NewClusterController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Cluster"}, "clusters", false, c.controllerFactory)
}
func (c *version) ClusterAlert() ClusterAlertController {
	return NewClusterAlertController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlert"}, "clusteralerts", true, c.controllerFactory)
}
func (c *version) ClusterAlertGroup() ClusterAlertGroupController {
	return NewClusterAlertGroupController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlertGroup"}, "clusteralertgroups", true, c.controllerFactory)
}
func (c *version) ClusterAlertRule() ClusterAlertRuleController {
	return NewClusterAlertRuleController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlertRule"}, "clusteralertrules", true, c.controllerFactory)
}
func (c *version) ClusterCatalog() ClusterCatalogController {
	return NewClusterCatalogController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterCatalog"}, "clustercatalogs", true, c.controllerFactory)
}
func (c *version) ClusterLogging() ClusterLoggingController {
	return NewClusterLoggingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterLogging"}, "clusterloggings", true, c.controllerFactory)
}
func (c *version) ClusterMonitorGraph() ClusterMonitorGraphController {
	return NewClusterMonitorGraphController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterMonitorGraph"}, "clustermonitorgraphs", true, c.controllerFactory)
}
func (c *version) ClusterRegistrationToken() ClusterRegistrationTokenController {
	return NewClusterRegistrationTokenController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterRegistrationToken"}, "clusterregistrationtokens", true, c.controllerFactory)
}
func (c *version) ClusterRoleTemplateBinding() ClusterRoleTemplateBindingController {
	return NewClusterRoleTemplateBindingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterRoleTemplateBinding"}, "clusterroletemplatebindings", true, c.controllerFactory)
}
func (c *version) ClusterScan() ClusterScanController {
	return NewClusterScanController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterScan"}, "clusterscans", true, c.controllerFactory)
}
func (c *version) ClusterTemplate() ClusterTemplateController {
	return NewClusterTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterTemplate"}, "clustertemplates", true, c.controllerFactory)
}
func (c *version) ClusterTemplateRevision() ClusterTemplateRevisionController {
	return NewClusterTemplateRevisionController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterTemplateRevision"}, "clustertemplaterevisions", true, c.controllerFactory)
}
func (c *version) ComposeConfig() ComposeConfigController {
	return NewComposeConfigController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ComposeConfig"}, "composeconfigs", false, c.controllerFactory)
}
func (c *version) DynamicSchema() DynamicSchemaController {
	return NewDynamicSchemaController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "DynamicSchema"}, "dynamicschemas", false, c.controllerFactory)
}
func (c *version) EtcdBackup() EtcdBackupController {
	return NewEtcdBackupController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "EtcdBackup"}, "etcdbackups", true, c.controllerFactory)
}
func (c *version) Feature() FeatureController {
	return NewFeatureController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Feature"}, "features", false, c.controllerFactory)
}
func (c *version) FleetWorkspace() FleetWorkspaceController {
	return NewFleetWorkspaceController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "FleetWorkspace"}, "fleetworkspaces", false, c.controllerFactory)
}
func (c *version) FreeIpaProvider() FreeIpaProviderController {
	return NewFreeIpaProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "FreeIpaProvider"}, "freeipaproviders", false, c.controllerFactory)
}
func (c *version) GithubProvider() GithubProviderController {
	return NewGithubProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GithubProvider"}, "githubproviders", false, c.controllerFactory)
}
func (c *version) GlobalDns() GlobalDnsController {
	return NewGlobalDnsController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalDns"}, "globaldnses", true, c.controllerFactory)
}
func (c *version) GlobalDnsProvider() GlobalDnsProviderController {
	return NewGlobalDnsProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalDnsProvider"}, "globaldnsproviders", true, c.controllerFactory)
}
func (c *version) GlobalRole() GlobalRoleController {
	return NewGlobalRoleController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalRole"}, "globalroles", false, c.controllerFactory)
}
func (c *version) GlobalRoleBinding() GlobalRoleBindingController {
	return NewGlobalRoleBindingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalRoleBinding"}, "globalrolebindings", false, c.controllerFactory)
}
func (c *version) GoogleOAuthProvider() GoogleOAuthProviderController {
	return NewGoogleOAuthProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GoogleOAuthProvider"}, "googleoauthproviders", false, c.controllerFactory)
}
func (c *version) Group() GroupController {
	return NewGroupController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Group"}, "groups", false, c.controllerFactory)
}
func (c *version) GroupMember() GroupMemberController {
	return NewGroupMemberController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GroupMember"}, "groupmembers", false, c.controllerFactory)
}
func (c *version) KontainerDriver() KontainerDriverController {
	return NewKontainerDriverController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "KontainerDriver"}, "kontainerdrivers", false, c.controllerFactory)
}
func (c *version) LocalProvider() LocalProviderController {
	return NewLocalProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "LocalProvider"}, "localproviders", false, c.controllerFactory)
}
func (c *version) MonitorMetric() MonitorMetricController {
	return NewMonitorMetricController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MonitorMetric"}, "monitormetrics", true, c.controllerFactory)
}
func (c *version) MultiClusterApp() MultiClusterAppController {
	return NewMultiClusterAppController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MultiClusterApp"}, "multiclusterapps", true, c.controllerFactory)
}
func (c *version) MultiClusterAppRevision() MultiClusterAppRevisionController {
	return NewMultiClusterAppRevisionController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "MultiClusterAppRevision"}, "multiclusterapprevisions", true, c.controllerFactory)
}
func (c *version) Node() NodeController {
	return NewNodeController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Node"}, "nodes", true, c.controllerFactory)
}
func (c *version) NodeDriver() NodeDriverController {
	return NewNodeDriverController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodeDriver"}, "nodedrivers", false, c.controllerFactory)
}
func (c *version) NodePool() NodePoolController {
	return NewNodePoolController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodePool"}, "nodepools", true, c.controllerFactory)
}
func (c *version) NodeTemplate() NodeTemplateController {
	return NewNodeTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodeTemplate"}, "nodetemplates", true, c.controllerFactory)
}
func (c *version) Notifier() NotifierController {
	return NewNotifierController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Notifier"}, "notifiers", true, c.controllerFactory)
}
func (c *version) OpenLdapProvider() OpenLdapProviderController {
	return NewOpenLdapProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "OpenLdapProvider"}, "openldapproviders", false, c.controllerFactory)
}
func (c *version) PodSecurityPolicyTemplate() PodSecurityPolicyTemplateController {
	return NewPodSecurityPolicyTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PodSecurityPolicyTemplate"}, "podsecuritypolicytemplates", false, c.controllerFactory)
}
func (c *version) PodSecurityPolicyTemplateProjectBinding() PodSecurityPolicyTemplateProjectBindingController {
	return NewPodSecurityPolicyTemplateProjectBindingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PodSecurityPolicyTemplateProjectBinding"}, "podsecuritypolicytemplateprojectbindings", true, c.controllerFactory)
}
func (c *version) Preference() PreferenceController {
	return NewPreferenceController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Preference"}, "preferences", true, c.controllerFactory)
}
func (c *version) Principal() PrincipalController {
	return NewPrincipalController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Principal"}, "principals", false, c.controllerFactory)
}
func (c *version) Project() ProjectController {
	return NewProjectController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Project"}, "projects", true, c.controllerFactory)
}
func (c *version) ProjectAlert() ProjectAlertController {
	return NewProjectAlertController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectAlert"}, "projectalerts", true, c.controllerFactory)
}
func (c *version) ProjectAlertGroup() ProjectAlertGroupController {
	return NewProjectAlertGroupController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectAlertGroup"}, "projectalertgroups", true, c.controllerFactory)
}
func (c *version) ProjectAlertRule() ProjectAlertRuleController {
	return NewProjectAlertRuleController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectAlertRule"}, "projectalertrules", true, c.controllerFactory)
}
func (c *version) ProjectCatalog() ProjectCatalogController {
	return NewProjectCatalogController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectCatalog"}, "projectcatalogs", true, c.controllerFactory)
}
func (c *version) ProjectLogging() ProjectLoggingController {
	return NewProjectLoggingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectLogging"}, "projectloggings", true, c.controllerFactory)
}
func (c *version) ProjectMonitorGraph() ProjectMonitorGraphController {
	return NewProjectMonitorGraphController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectMonitorGraph"}, "projectmonitorgraphs", true, c.controllerFactory)
}
func (c *version) ProjectNetworkPolicy() ProjectNetworkPolicyController {
	return NewProjectNetworkPolicyController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectNetworkPolicy"}, "projectnetworkpolicies", true, c.controllerFactory)
}
func (c *version) ProjectRoleTemplateBinding() ProjectRoleTemplateBindingController {
	return NewProjectRoleTemplateBindingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectRoleTemplateBinding"}, "projectroletemplatebindings", true, c.controllerFactory)
}
func (c *version) RkeAddon() RkeAddonController {
	return NewRkeAddonController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeAddon"}, "rkeaddons", true, c.controllerFactory)
}
func (c *version) RkeK8sServiceOption() RkeK8sServiceOptionController {
	return NewRkeK8sServiceOptionController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeK8sServiceOption"}, "rkek8sserviceoptions", true, c.controllerFactory)
}
func (c *version) RkeK8sSystemImage() RkeK8sSystemImageController {
	return NewRkeK8sSystemImageController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeK8sSystemImage"}, "rkek8ssystemimages", true, c.controllerFactory)
}
func (c *version) RoleTemplate() RoleTemplateController {
	return NewRoleTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RoleTemplate"}, "roletemplates", false, c.controllerFactory)
}
func (c *version) SamlProvider() SamlProviderController {
	return NewSamlProviderController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "SamlProvider"}, "samlproviders", false, c.controllerFactory)
}
func (c *version) SamlToken() SamlTokenController {
	return NewSamlTokenController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "SamlToken"}, "samltokens", false, c.controllerFactory)
}
func (c *version) Setting() SettingController {
	return NewSettingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Setting"}, "settings", false, c.controllerFactory)
}
func (c *version) Template() TemplateController {
	return NewTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Template"}, "templates", false, c.controllerFactory)
}
func (c *version) TemplateContent() TemplateContentController {
	return NewTemplateContentController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "TemplateContent"}, "templatecontents", false, c.controllerFactory)
}
func (c *version) TemplateVersion() TemplateVersionController {
	return NewTemplateVersionController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "TemplateVersion"}, "templateversions", false, c.controllerFactory)
}
func (c *version) Token() TokenController {
	return NewTokenController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Token"}, "tokens", false, c.controllerFactory)
}
func (c *version) User() UserController {
	return NewUserController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "User"}, "users", false, c.controllerFactory)
}
func (c *version) UserAttribute() UserAttributeController {
	return NewUserAttributeController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "UserAttribute"}, "userattributes", false, c.controllerFactory)
}
