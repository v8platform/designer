package designer

import (
	"context"
	"github.com/v8platform/errors"
	"github.com/v8platform/marshaler"
	"net"
	"strconv"
	"time"
)

// AgentModeOptions Включает режим агента конфигуратора.
//	При наличии этой команды игнорируются команды /DisableStartupMessages /DisableStartupDialogs, если таковые указаны.
//
type AgentModeOptions struct {
	command struct{} `v8:"/AgentMode" json:"-"`

	///AgentBaseDir <рабочий каталог>
	//Данная команда позволяет указать рабочий каталог,
	//который используется при работе SFTP-сервера, а также при работе команд загрузки/выгрузки конфигурации.
	//Если команда не указана, то будет использован следующий каталог:
	//	Для ОС Windows: %LOCALAPPDATA%\1C\1cv8\<Уникальный идентификатор информационной базы>\sftp.
	//	Для ОС Linux: ~/.1cv8/1C/1cv8/<Уникальный идентификатор информационной базы>/sftp.
	//	Для ОС macOS: ~/.1cv8/1C/1cv8/<Уникальный идентификатор информационной базы>/sftp.
	BaseDir string `v8:"/AgentBaseDir, optional" json:"dir"`

	///AgentPort <Порт>
	//Указывает номер TCP-порта, который использует агент в режиме SSH-сервера.
	//Если команда не указана, то по умолчанию используется TCP-порт с номером 1543.
	Port int `v8:"/AgentPort, optional" json:"port"`

	///AgentListenAddress <Адрес>
	//Параметр команды позволяет указать IP-адрес, который будет прослушиваться агентом.
	//Если команда не указан, то по умолчанию используется IP-адрес 127.0.0.1.
	ListenAddress string `v8:"/AgentListenAddress, optional" json:"ip"`

	// SSHHostKeyAuto Команда указывает, что закрытый ключ хоста имеет следующее расположение (в зависимости от используемой операционной системы):
	//	Для ОС Windows: %LOCALAPPDATA%\1C\1cv8\host_id.
	//	Для ОС Linux: ~/.1cv8/1C/1cv8/host_id.
	//	Для ОС macOS: ~/.1cv8/1C/1cv8/host_id.
	//	Если указанный файл не будет обнаружен, то будет создан закрытый ключ для алгоритма RSA с длиной ключа 2 048 бит.
	SSHHostKeyAuto bool `v8:"/AgentSSHHostKeyAuto, optional" json:"ssh-auto"`

	///AgentSSHHostKey <приватный ключ>
	//Параметр команды позволяет указать путь к закрытому ключу хоста.
	//Если параметр не указан, то должна быть указана команда /AgentSSHHostKeyAuto.
	//Если не указан ни одна команда ‑ запуск в режиме агента будет невозможен.
	SSHHostKey string `v8:"/AgentSSHHostKey, optional" json:"ssh-key"`

	Visible bool `v8:"/Visible" json:"visible"`
}

func (d AgentModeOptions) Command() string {
	return COMMAND_DESIGNER
}

func (d AgentModeOptions) Check() error {

	if !d.SSHHostKeyAuto && len(d.SSHHostKey) == 0 {

		return errors.Check.New("ssh host key must be set").WithContext("msg", "field SSHHostKeyAuto or SSHHostKey not set")

	}

	return nil
}

func (d AgentModeOptions) Values() []string {

	v, _ := marshaler.Marshal(d)
	return v
}

func (o AgentModeOptions) WithBaseDir(dir string) AgentModeOptions {

	newO := o
	newO.BaseDir = dir
	return newO

}

func (o AgentModeOptions) WithListenAddress(ipPort string) AgentModeOptions {

	host, portString, _ := net.SplitHostPort(ipPort)

	port, _ := strconv.ParseInt(portString, 10, 64)

	newO := o
	newO.ListenAddress = host
	newO.Port = int(port)
	return newO

}

func waitAgent(ctx context.Context, hostPort string) error {

	ready := make(chan error)

	go func() {
		timeuot, _ := context.WithTimeout(ctx, time.Second*10)
		ticker := time.Tick(time.Second)
		for {
			select {
			case <-ready:
				return
			case <-ticker:

				_, err := net.Dial("tcp", hostPort)
				if err == nil {
					close(ready)
					return
				}

			case <-timeuot.Done():
				ready <- timeuot.Err()
			}
		}

	}()
	err := <-ready

	return err

}

func (o AgentModeOptions) Wait(ctx context.Context) error {

	return waitAgent(ctx, net.JoinHostPort(o.ListenAddress, strconv.FormatInt(int64(o.Port), 64)))

}
