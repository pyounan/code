// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"os"
	"testing"
)

func TestHelloPubSub(t *testing.T) {
	os.Setenv("MAIL_ENABLED", "true")
	os.Setenv("SLACK_ENABLED", "true")
	os.Setenv("SLACK_FILTER_BUILD_STATUSES", "")

	os.Setenv("SLACK_WEBHOOK_URL", "<WEBHOOK>")
	os.Setenv("MAIL_DOMAIN", "notify.tonystarkapi.com")
	os.Setenv("MAIL_SENDER", "notify@notify.tonystarkapi.com")
	os.Setenv("MAIL_APIKEY", "<API-KEY>")
	os.Setenv("RECIPIENTS", "xmarcoied@gmail.com")

	data2 := []byte(`{
		"name": "projects/323324163496/locations/global/builds/2b4d61d6-8ded-4395-8c5b-4fcfbdca610d",
		"id": "2b4d61d6-8ded-4395-8c5b-4fcfbdca610d",
		"projectId": "es-webkings",
		"status": "SUCCESS",
		"source": {
			"storageSource": {
				"bucket": "323324163496.cloudbuild-source.googleusercontent.com",
				"object": "53fd63bea67d222bd624e58bf6702d5b75278f65-b85c5ee2-432d-46b3-8169-b43867ec4568.tar.gz"
			}
		},
		"steps": [
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"clean",
					"compile",
					"install",
					"-DskipTests"
				],
				"id": "maven-compile-install",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:34:35.760184400Z",
					"endTime": "2020-10-27T09:38:37.989228664Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:34:35.760184400Z",
					"endTime": "2020-10-27T09:34:42.219307796Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/zuul:master-53fd63b"
				],
				"dir": "externals/zuul",
				"id": "maven-jib-zuul",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:38:37.989335780Z",
					"endTime": "2020-10-27T09:38:50.832866367Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:38:37.989335780Z",
					"endTime": "2020-10-27T09:38:37.990711196Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/wallet-service:master-53fd63b"
				],
				"dir": "services/wallet-service",
				"id": "maven-jib-wallet-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:38:50.832990072Z",
					"endTime": "2020-10-27T09:39:14.914452313Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:38:50.832990072Z",
					"endTime": "2020-10-27T09:38:50.834252999Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/events-service:master-53fd63b"
				],
				"dir": "services/events-service",
				"id": "maven-jib-events-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:39:14.914569633Z",
					"endTime": "2020-10-27T09:39:26.459519060Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:39:14.914569633Z",
					"endTime": "2020-10-27T09:39:14.916126726Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/accounts-service:master-53fd63b"
				],
				"dir": "services/accounts-service",
				"id": "maven-jib-accounts-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:39:26.459625486Z",
					"endTime": "2020-10-27T09:39:40.840699890Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:39:26.459625486Z",
					"endTime": "2020-10-27T09:39:26.461037732Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/board-service:master-53fd63b"
				],
				"dir": "services/board-service",
				"id": "maven-jib-board-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:39:40.840814255Z",
					"endTime": "2020-10-27T09:39:52.157097883Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:39:40.840814255Z",
					"endTime": "2020-10-27T09:39:40.842290401Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/questions-service:master-53fd63b"
				],
				"dir": "services/questions-service",
				"id": "maven-jib-questions-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:39:52.157236815Z",
					"endTime": "2020-10-27T09:40:04.711542582Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:39:52.157236815Z",
					"endTime": "2020-10-27T09:39:52.158878597Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/info-service:master-53fd63b"
				],
				"dir": "services/info-service",
				"id": "maven-jib-info-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:40:04.711668592Z",
					"endTime": "2020-10-27T09:40:18.657965275Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:40:04.711668592Z",
					"endTime": "2020-10-27T09:40:04.713229666Z"
				},
				"status": "SUCCESS"
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/notifications-service:master-53fd63b"
				],
				"dir": "services/notifications-service",
				"id": "maven-jib-notifications-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				],
				"timing": {
					"startTime": "2020-10-27T09:40:18.658093707Z",
					"endTime": "2020-10-27T09:40:32.661810199Z"
				},
				"pullTiming": {
					"startTime": "2020-10-27T09:40:18.658093707Z",
					"endTime": "2020-10-27T09:40:18.659448179Z"
				},
				"status": "SUCCESS"
			}
		],
		"results": {
			"buildStepImages": [
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				""
			],
			"buildStepOutputs": [
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				""
			]
		},
		"createTime": "2020-10-27T09:33:36.588802999Z",
		"startTime": "2020-10-27T09:34:25.527055121Z",
		"finishTime": "2020-10-27T09:40:33.337406Z",
		"timeout": "600s",
		"queueTtl": "3600s",
		"logsBucket": "gs://323324163496.cloudbuild-logs.googleusercontent.com",
		"sourceProvenance": {
			"resolvedStorageSource": {
				"bucket": "323324163496.cloudbuild-source.googleusercontent.com",
				"object": "53fd63bea67d222bd624e58bf6702d5b75278f65-b85c5ee2-432d-46b3-8169-b43867ec4568.tar.gz",
				"generation": "1603791215995479"
			},
			"fileHashes": {
				"gs://323324163496.cloudbuild-source.googleusercontent.com/53fd63bea67d222bd624e58bf6702d5b75278f65-b85c5ee2-432d-46b3-8169-b43867ec4568.tar.gz#1603791215995479": {
					"fileHash": [
						{
							"type": "MD5",
							"value": "/brhHhvHDhglahvw/jZOiA=="
						}
					]
				}
			}
		},
		"buildTriggerId": "f9de2b60-3dc5-4b09-85f9-61d2aafd1a22",
		"options": {
			"machineType": "N1_HIGHCPU_8",
			"substitutionOption": "ALLOW_LOOSE",
			"dynamicSubstitutions": true,
			"logging": "LEGACY"
		},
		"logUrl": "https://console.cloud.google.com/cloud-build/builds/2b4d61d6-8ded-4395-8c5b-4fcfbdca610d?project=323324163496",
		"substitutions": {
			"BRANCH_NAME": "master",
			"COMMIT_SHA": "53fd63bea67d222bd624e58bf6702d5b75278f65",
			"REPO_NAME": "auc-backend",
			"REVISION_ID": "53fd63bea67d222bd624e58bf6702d5b75278f65",
			"SHORT_SHA": "53fd63b",
			"_SERVICE_DIR": "externals/zuul",
			"_SERVICE_NAME": "zuul"
		},
		"tags": [
			"trigger-f9de2b60-3dc5-4b09-85f9-61d2aafd1a22"
		],
		"timing": {
			"BUILD": {
				"startTime": "2020-10-27T09:34:35.077742471Z",
				"endTime": "2020-10-27T09:40:32.661928341Z"
			},
			"FETCHSOURCE": {
				"startTime": "2020-10-27T09:34:26.139916379Z",
				"endTime": "2020-10-27T09:34:35.077658722Z"
			}
		}
	}`)
	data1 := []byte(`{
		"name": "projects/323324163496/locations/global/builds/21da73ea-4ff1-433b-9d32-8c5ad87a0829",
		"id": "21da73ea-4ff1-433b-9d32-8c5ad87a0829",
		"projectId": "es-webkings",
		"status": "SUCCESS",
		"source": {
			"storageSource": {
				"bucket": "323324163496.cloudbuild-source.googleusercontent.com",
				"object": "53fd63bea67d222bd624e58bf6702d5b75278f65-b4837b4e-c0bd-4e6d-bab3-963136226dc6.tar.gz"
			}
		},
		"steps": [
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"clean",
					"compile",
					"install",
					"-DskipTests"
				],
				"id": "maven-compile-install",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/zuul:master-53fd63b"
				],
				"dir": "externals/zuul",
				"id": "maven-jib-zuul",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/wallet-service:master-53fd63b"
				],
				"dir": "services/wallet-service",
				"id": "maven-jib-wallet-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/events-service:master-53fd63b"
				],
				"dir": "services/events-service",
				"id": "maven-jib-events-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/accounts-service:master-53fd63b"
				],
				"dir": "services/accounts-service",
				"id": "maven-jib-accounts-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/board-service:master-53fd63b"
				],
				"dir": "services/board-service",
				"id": "maven-jib-board-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/questions-service:master-53fd63b"
				],
				"dir": "services/questions-service",
				"id": "maven-jib-questions-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/info-service:master-53fd63b"
				],
				"dir": "services/info-service",
				"id": "maven-jib-info-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			},
			{
				"name": "maven:3.6.3-jdk-8",
				"args": [
					"compile",
					"com.google.cloud.tools:jib-maven-plugin:2.0.0:build",
					"-Dimage=gcr.io/es-webkings/github.com/extremesolution/auc-backend/notifications-service:master-53fd63b"
				],
				"dir": "services/notifications-service",
				"id": "maven-jib-notifications-service",
				"entrypoint": "mvn",
				"volumes": [
					{
						"name": "mavenrepo",
						"path": "/root/.m2"
					}
				]
			}
		],
		"createTime": "2020-10-27T09:29:56.856744071Z",
		"timeout": "600s",
		"queueTtl": "3600s",
		"logsBucket": "gs://323324163496.cloudbuild-logs.googleusercontent.com",
		"sourceProvenance": {
			"resolvedStorageSource": {
				"bucket": "323324163496.cloudbuild-source.googleusercontent.com",
				"object": "53fd63bea67d222bd624e58bf6702d5b75278f65-b4837b4e-c0bd-4e6d-bab3-963136226dc6.tar.gz",
				"generation": "1603790996096166"
			}
		},
		"buildTriggerId": "f9de2b60-3dc5-4b09-85f9-61d2aafd1a22",
		"options": {
			"machineType": "N1_HIGHCPU_8",
			"substitutionOption": "ALLOW_LOOSE",
			"dynamicSubstitutions": true,
			"logging": "LEGACY"
		},
		"logUrl": "https://console.cloud.google.com/cloud-build/builds/21da73ea-4ff1-433b-9d32-8c5ad87a0829?project=323324163496",
		"substitutions": {
			"BRANCH_NAME": "master",
			"COMMIT_SHA": "53fd63bea67d222bd624e58bf6702d5b75278f65",
			"REPO_NAME": "auc-backend",
			"REVISION_ID": "53fd63bea67d222bd624e58bf6702d5b75278f65",
			"SHORT_SHA": "53fd63b",
			"_SERVICE_DIR": "externals/zuul",
			"_SERVICE_NAME": "zuul"
		},
		"tags": [
			"trigger-f9de2b60-3dc5-4b09-85f9-61d2aafd1a22"
		]
	}`)

	_ = data1
	type args struct {
		ctx context.Context
		m   PubSubMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Notification | Slack",
			args: args{
				ctx: context.Background(),
				m: PubSubMessage{
					Data: data2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HelloPubSub(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("HelloPubSub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
