package teamcity

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/icelander/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestWorkingDirectory(t *testing.T) {
         wd, _ := os.Getwd()
         t.Log(wd)
}

func TestClientGetBuildProperties(t *testing.T) {
	client := NewTestClient(newResponse(`{"property":[{"name": "build.counter", "value": "12"}], "count": 1}`), nil)

	props, err := client.GetBuildProperties("999999")

	if len(props) != 1 {
		t.Fatal("Expected to have 1 property, found", len(props))
	}

	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
}

func TestClientGetBuilds(t *testing.T) {
	client := NewTestClient(newResponse(`{"count":2,"build":[{"id":2,"taskId":2,"buildTypeId":"MattermostTeamcityPlugin_Build","buildTypeInternalId":"bt2","number":"2","status":"SUCCESS","state":"finished","running":false,"composite":false,"failedToStart":false,"personal":false,"history":false,"pinned":false,"href":"/app/rest/builds/id:2","webUrl":"http://127.0.0.1:8111/viewLog.html?buildId=2&buildTypeId=MattermostTeamcityPlugin_Build","limitedChangesCount":0,"artifactsDirectory":"/opt/TeamCity/.BuildServer/system/artifacts/MattermostTeamcityPlugin/Build/2","links":{"count":1,"link":[{"type":"webView","url":"http://127.0.0.1:8111/viewLog.html?buildId=2&buildTypeId=MattermostTeamcityPlugin_Build","relativeUrl":"/viewLog.html?buildId=2&buildTypeId=MattermostTeamcityPlugin_Build"}]},"statusText":"Success","buildType":{"id":"MattermostTeamcityPlugin_Build","name":"Build","projectName":"Mattermost Teamcity Plugin","projectId":"MattermostTeamcityPlugin","href":"/app/rest/buildTypes/id:MattermostTeamcityPlugin_Build","webUrl":"http://127.0.0.1:8111/viewType.html?buildTypeId=MattermostTeamcityPlugin_Build"},"tags":{"tag":[]},"queuedDate":"20200119T190211+0000","startDate":"20200119T190249+0000","finishDate":"20200119T190252+0000","triggered":{"type":"user","date":"20200119T190211+0000","displayText":"you","rawValue":"##userId='1' type='user'","user":{"username":"paul","name":"Paul Rothrock","id":1,"href":"/app/rest/users/id:1"},"properties":{"count":2}},"changes":{"count":0,"href":"/app/rest/changes?locator=build:(id:2)","change":[]},"revisions":{"count":1,"revision":[{"version":"2691bc37fefa5216ace02434b8a24d042013bea9","vcsBranchName":"refs/heads/master","vcs-root-instance":{}}]},"artifactDependencyChanges":{"count":0},"agent":{"id":1,"name":"ip_10.0.2.15","typeId":1,"href":"/app/rest/agents/id:1","webUrl":"http://127.0.0.1:8111/agentDetails.html?id=1&agentTypeId=1&realAgentName=ip_10.0.2.15"},"artifacts":{"count":0,"href":"/app/rest/builds/id:2/artifacts/children/"},"relatedIssues":{"href":"/app/rest/builds/id:2/relatedIssues"},"properties":{"property":[]},"resultingProperties":{"count":136},"attributes":{"count":0},"statistics":{"href":"/app/rest/builds/id:2/statistics"},"metadata":{"count":0,"data":[]},"snapshot-dependencies":{"count":0},"artifact-dependencies":{"count":0},"settingsHash":"a54b485f6762fa2f4185b2f5206f5ac2d0fed2d7","currentSettingsHash":"a54b485f6762fa2f4185b2f5206f5ac2d0fed2d7","modificationId":"-1","chainModificationId":"-1","replacementIds":{"item":[]},"related":{},"usedByOtherBuilds":false},{"id":1,"taskId":1,"buildTypeId":"MattermostTeamcityPlugin_Build","buildTypeInternalId":"bt2","number":"1","status":"SUCCESS","state":"finished","running":false,"composite":false,"failedToStart":false,"personal":false,"history":false,"pinned":false,"href":"/app/rest/builds/id:1","webUrl":"http://127.0.0.1:8111/viewLog.html?buildId=1&buildTypeId=MattermostTeamcityPlugin_Build","limitedChangesCount":0,"artifactsDirectory":"/opt/TeamCity/.BuildServer/system/artifacts/MattermostTeamcityPlugin/Build/1","links":{"count":1,"link":[{"type":"webView","url":"http://127.0.0.1:8111/viewLog.html?buildId=1&buildTypeId=MattermostTeamcityPlugin_Build","relativeUrl":"/viewLog.html?buildId=1&buildTypeId=MattermostTeamcityPlugin_Build"}]},"statusText":"Success","buildType":{"id":"MattermostTeamcityPlugin_Build","name":"Build","projectName":"Mattermost Teamcity Plugin","projectId":"MattermostTeamcityPlugin","href":"/app/rest/buildTypes/id:MattermostTeamcityPlugin_Build","webUrl":"http://127.0.0.1:8111/viewType.html?buildTypeId=MattermostTeamcityPlugin_Build"},"tags":{"tag":[]},"queuedDate":"20200118T213458+0000","startDate":"20200118T214343+0000","finishDate":"20200118T214349+0000","triggered":{"type":"user","date":"20200118T213458+0000","displayText":"Super user","rawValue":"##userId='-42' type='user'","user":{"username":"","name":"Super user","id":-42,"href":"/app/rest/users/id:-42"},"properties":{"count":2}},"changes":{"count":0,"href":"/app/rest/changes?locator=build:(id:1)","change":[]},"revisions":{"count":1,"revision":[{"version":"2691bc37fefa5216ace02434b8a24d042013bea9","vcsBranchName":"refs/heads/master","vcs-root-instance":{}}]},"artifactDependencyChanges":{"count":0},"agent":{"id":1,"name":"ip_10.0.2.15","typeId":1,"href":"/app/rest/agents/id:1","webUrl":"http://127.0.0.1:8111/agentDetails.html?id=1&agentTypeId=1&realAgentName=ip_10.0.2.15"},"artifacts":{"count":0,"href":"/app/rest/builds/id:1/artifacts/children/"},"relatedIssues":{"href":"/app/rest/builds/id:1/relatedIssues"},"properties":{"property":[]},"resultingProperties":{"count":136},"attributes":{"count":0},"statistics":{"href":"/app/rest/builds/id:1/statistics"},"metadata":{"count":0,"data":[]},"snapshot-dependencies":{"count":0},"artifact-dependencies":{"count":0},"settingsHash":"24f2ce2d8821500cc6c5a1759235566a5462b03e","currentSettingsHash":"24f2ce2d8821500cc6c5a1759235566a5462b03e","modificationId":"-1","chainModificationId":"-1","replacementIds":{"item":[]},"related":{},"usedByOtherBuilds":false}]}`), nil)

	builds, err := client.GetBuilds()

	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	if len(builds) != 2 {
		t.Fatal("Expected to have two builds, found ", len(builds))
	}
}

func TestClientCancelBuild(t *testing.T) {
	assert := assert.New(t)

	client := NewTestClient(newResponse(``), nil)

	build, err := client.CancelBuild(1234, "comment string")

	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	assert.IsType(build, &types.Build{})
	assert.Equal(build.ID, int64(29))
}

func TestClientGetAgentStats(t *testing.T) {
	assert := assert.New(t)

	jsonFixture, err := ioutil.ReadFile("../fixtures/TestClientGetAgentStats.json")

	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	client := NewTestClient(newResponse(string(jsonFixture)), nil)

	agents, err := client.GetAgentStats()
	
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	assert.Equal(len(agents), 1)
	assert.IsType(agents[0], &types.Agent{})

	assert.Equal(agents[0].ID, 1)
	assert.Equal(agents[0].Name, "ip_10.0.2.15")
	assert.Equal(agents[0].ActiveBuild.BuildTypeID, "MattermostTeamcityPlugin_TestBuild")
	assert.True(agents[0].Connected)
}

func TestClientGetBuildQueue(t *testing.T) {
	assert := assert.New(t)

	jsonFixture, err := ioutil.ReadFile("../fixtures/TestClientGetBuildQueue.json")

	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	client := NewTestClient(newResponse(string(jsonFixture)), nil)

	agents, err := client.GetBuildQueue()
	
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
}