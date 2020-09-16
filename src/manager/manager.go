package manager

import "log"

type (
	AppManager struct {
		appShootdown chan interface{}
		appReload    chan interface{}
		errors       chan error
	}
)

func (mgr *AppManager) Start() {
	mgr.appReload = make(chan interface{})
	mgr.appShootdown = make(chan interface{})
	mgr.errors = make(chan error)
	for {
		select {
		case err := <-mgr.errors:
			log.Println(err, "чет все сложилось нада чет делать")
		case <-mgr.appShootdown:
			// делаем все что нужн если упало
			//выводим ошибку и роняем все к х*ям
			//берем идентификатор гуя из контекста
			// апку додумать как убить
		case <-mgr.appReload:
			//хз нах*я вдруг гуй отвалится перезапустить его
		}
	}

}
func (mgr *AppManager) HandleError(err error) {
	//обрабатываю ошибку которую поймал
	// пхаю все в кАнал
	mgr.errors <- err
}
