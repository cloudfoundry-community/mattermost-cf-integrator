package mci_test

import (
	. "github.com/cloudfoundry-community/mattermost-cf-integrator/mci"
	"path"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Mci", func() {
	var vcapApplication, vcapServices, port, dir, fixturePath string
	var err error
	dir, err = os.Getwd()
	if err != nil {
		Fail(err.Error())
	}
	fixturePath = path.Join(dir, "..", "fixtures")
	configPath := path.Join(fixturePath, "config.json")
	BeforeEach(func() {
		vcapApplication = os.Getenv("VCAP_APPLICATION")
		vcapServices = os.Getenv("VCAP_SERVICES")
		port = os.Getenv("PORT")
	})
	AfterEach(func() {
		os.Setenv("VCAP_APPLICATION", vcapApplication)
		os.Setenv("VCAP_SERVICES", vcapServices)
		os.Setenv("PORT", port)
	})
	Context("not in cloud foundry", func() {
		It("should not start anything", func() {
			os.Setenv("VCAP_APPLICATION", "")
			Expect(IsInCloudFoundry()).To(BeFalse())
			err := CloudifyConfig(&MattermostConfig{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Not in Cloud Foundry"))
		})
	})
	Context("in cloud foundry", func() {
		Describe("when database service is mysql", func() {
			It("should update config file for mattermost to consider right port and mysql service found by tag", func() {
				var expectedMattermostConfig *MattermostConfig
				expectedConfigPath := path.Join(fixturePath, "config-mysql.json")
				expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
				Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

				config, err := ExtractConfig(configPath)
				Expect(err).NotTo(HaveOccurred())
				err = CloudifyConfig(config)
				Expect(err).NotTo(HaveOccurred())
				Expect(config).To(Equal(expectedMattermostConfig))
			})
			It("should update config file for mattermost to consider right port and mysql service found by service name", func() {
				var expectedMattermostConfig *MattermostConfig
				os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"mysql://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[]}]}")
				expectedConfigPath := path.Join(fixturePath, "config-mysql.json")
				expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
				Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

				config, err := ExtractConfig(configPath)
				Expect(err).NotTo(HaveOccurred())
				err = CloudifyConfig(config)
				Expect(err).NotTo(HaveOccurred())
				Expect(config).To(Equal(expectedMattermostConfig))
			})
		})

		Describe("when database service is postgres", func() {
			It("should update config file for mattermost to consider right port and postgres service found by tag", func() {
				var expectedMattermostConfig *MattermostConfig
				os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"postgres://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[\"postgres\"]}]}")
				expectedConfigPath := path.Join(fixturePath, "config-postgres.json")
				expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
				Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

				config, err := ExtractConfig(configPath)
				Expect(err).NotTo(HaveOccurred())
				err = CloudifyConfig(config)
				Expect(err).NotTo(HaveOccurred())
				Expect(config).To(Equal(expectedMattermostConfig))
			})
			It("should update config file for mattermost to consider right port and postgres service found by name", func() {
				var expectedMattermostConfig *MattermostConfig
				os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"postgres://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"postgres-mattermost\",\"plan\":\"spark\",\"tags\":[]}]}")
				expectedConfigPath := path.Join(fixturePath, "config-postgres.json")
				expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
				Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

				config, err := ExtractConfig(configPath)
				Expect(err).NotTo(HaveOccurred())
				err = CloudifyConfig(config)
				Expect(err).NotTo(HaveOccurred())
				Expect(config).To(Equal(expectedMattermostConfig))
			})
		})
		Describe("when have a postgres and a mysql services", func() {
			It("should choose postgres database", func() {
				os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"postgres://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[\"mysql\"]}],\"postgresdb\":[{\"credentials\":{\"uri\":\"postgres://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[\"postgres\"]}]}")
				var expectedMattermostConfig *MattermostConfig
				expectedConfigPath := path.Join(fixturePath, "config-postgres.json")
				expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
				Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

				config, err := ExtractConfig(configPath)
				Expect(err).NotTo(HaveOccurred())
				err = CloudifyConfig(config)
				Expect(err).NotTo(HaveOccurred())
				Expect(config).To(Equal(expectedMattermostConfig))
			})
		})
		Describe("when there is no database service", func() {
			It("should complain that's there is no database found", func() {
				os.Setenv("VCAP_SERVICES", "{}")
				config, err := ExtractConfig(configPath)
				Expect(err).NotTo(HaveOccurred())
				err = CloudifyConfig(config)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Cannot find database"))
			})
		})
		Describe("when there a smtp service", func() {
			Context("with a sendgrid smtp provided", func() {
				It("should fill information about the sendgrid smtp server", func() {
					os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"mysql://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[]}], \"sendgrid\":[{\"credentials\":{\"hostname\":\"smtp.host.com\",\"password\":\"password\",\"username\":\"user\"},\"label\":\"sendgrid\",\"name\":\"test-sendgrid\",\"plan\":\"free\",\"provider\":null,\"syslog_drain_url\":null,\"tags\":[\"Retail\",\"Email\",\"smtp\",\"Inventorymanagement\"]}]}")
					var expectedMattermostConfig *MattermostConfig
					expectedConfigPath := path.Join(fixturePath, "config-smtp-sendgrid.json")
					expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
					Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

					config, err := ExtractConfig(configPath)
					Expect(err).NotTo(HaveOccurred())
					err = CloudifyConfig(config)
					Expect(err).NotTo(HaveOccurred())
					Expect(config).To(Equal(expectedMattermostConfig))
				})
			})
			Context("with an unknown smtp provided", func() {
				It("should fill information about a default smtp server", func() {
					os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"mysql://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[]}], \"unknown-smtp\":[{\"credentials\":{\"hostname\":\"smtp.host.com\",\"password\":\"password\",\"username\":\"user\"},\"label\":\"unknown-smtp\",\"name\":\"test-sendgrid\",\"plan\":\"free\",\"provider\":null,\"syslog_drain_url\":null,\"tags\":[\"Retail\",\"Email\",\"smtp\",\"Inventorymanagement\"]}]}")
					var expectedMattermostConfig *MattermostConfig
					expectedConfigPath := path.Join(fixturePath, "config-smtp-default.json")
					expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
					Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

					config, err := ExtractConfig(configPath)
					Expect(err).NotTo(HaveOccurred())
					err = CloudifyConfig(config)
					Expect(err).NotTo(HaveOccurred())
					Expect(config).To(Equal(expectedMattermostConfig))
				})
			})
		})
		Describe("when there a s3 service", func() {
			Context("with an s3 broker with uri", func() {
				It("should fill information about the s3 server", func() {
					os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"mysql://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[]}], \"p-riakcs\":[{\"credentials\":{\"uri\":\"https://BU8FRUIT:MGB8A%3D%3D@p-riakcs.myriak.com/service-instance\"},\"label\":\"p-riakcs\",\"name\":\"riak-service-db-dumper\",\"plan\":\"developer\",\"tags\":[\"riak-cs\",\"s3\"]}]}")
					var expectedMattermostConfig *MattermostConfig
					expectedConfigPath := path.Join(fixturePath, "config-s3.json")
					expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
					Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

					config, err := ExtractConfig(configPath)
					Expect(err).NotTo(HaveOccurred())
					err = CloudifyConfig(config)
					Expect(err).NotTo(HaveOccurred())
					Expect(config).To(Equal(expectedMattermostConfig))
				})
			})

			Context("with an cloudfoundry-community/s3-cf-service-broker", func() {
				It("should fill information about the s3 server", func() {
					os.Setenv("VCAP_SERVICES", "{\"cleardb\":[{\"credentials\":{\"uri\":\"mysql://titi:toto@my.db.com:3306/mydbname?reconnect=true\"},\"label\":\"cleardb\",\"name\":\"mysql-mattermost\",\"plan\":\"spark\",\"tags\":[]}], \"p-riakcs\":[{\"credentials\":{\"username\":\"a-user\",\"access_key_id\":\"BU8FRUIT\",\"bucket\":\"service-instance\",\"secret_access_key\":\"MGB8A==\"},\"label\":\"p-riakcs\",\"name\":\"riak-service-db-dumper\",\"plan\":\"developer\",\"tags\":[\"s3\"]}]}")
					var expectedMattermostConfig *MattermostConfig
					expectedConfigPath := path.Join(fixturePath, "config-s3-amazon-broker.json")
					expectedMattermostConfig, err := ExtractConfig(expectedConfigPath)
					Expect(err).NotTo(HaveOccurred(), "Problem during loading expected json")

					config, err := ExtractConfig(configPath)
					Expect(err).NotTo(HaveOccurred())
					err = CloudifyConfig(config)
					Expect(err).NotTo(HaveOccurred())
					Expect(config).To(Equal(expectedMattermostConfig))
				})
			})
		})
	})
})
