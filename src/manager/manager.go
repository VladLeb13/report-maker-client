package manager

import "log"

type (
	AppManager struct {
		isAppDown   ifdown
		isAppReload ifreload
		Errors      chan error
	}

	ifdown struct {
		appname string
		down    chan interface{}
	}
	ifreload struct {
		appname string
		reload  chan interface{}
	}
)

func (mgr *AppManager) Start() {
	for {
		select {
		case <-mgr.isAppDown.down:
			// делаем все что нужн если упало
			//выводим ошибку и роняем все к х*ям
			//берем идентификатор гуя из контекста
			// апку додумать как убить
		case <-mgr.isAppReload.reload:
			//хз нах*я вдруг гуй отвалится перезапустить его
		case err := <-mgr.Errors:
			log.Println(err, "чет все сложилось нада чет делать")
		}
	}

}
