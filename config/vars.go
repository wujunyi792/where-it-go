package config

type GlobalConfig struct {
	MODE        string
	ProgramName string
	AUTHOR      string
	VERSION     string
	Auth        struct {
		Secret string
		Issuer string
	}
	SQL struct {
		Use    bool
		Config struct {
			TYPE     string
			IP       string
			PORT     string
			USER     string
			PASSWORD string
			DATABASE string
		}
	}
	REDIS struct {
		Use    bool
		Config struct {
			IP       string
			PORT     string
			PASSWORD string
			DB       int
		}
	}
	OSS struct {
		Use    bool
		Config struct {
			AccessKeySecret string
			AccessKeyId     string
			EndPoint        string
			BucketName      string
			BaseURL         string
			Path            string
			CallbackUrl     string
			ExpireTime      int64
		}
	}
	Mail struct {
		Use    bool
		Config struct {
			SMTP     string
			PORT     int
			ACCOUNT  string
			PASSWORD string
		}
	}
	CMS struct {
		Use    bool
		Config struct {
			SecretId   string
			SecretKey  string
			AppId      string
			TemplateId string
			Sign       string
		}
	}
}
