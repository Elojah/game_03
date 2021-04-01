import {Scene} from "phaser";

export class Login extends Scene {

    HTMLlogin: Phaser.GameObjects.DOMElement;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    preload() {
        this.load.html('login', 'html/login.html')
        this.load.image('home_background_00', 'img/home_background_00.png')
    }
    create() {
        if (document.cookie.startsWith("oauth-session=")) {
            this.scene.transition({
                target: "home",
                duration: 1000,
                remove: true,
            })

            return
        }

        this.add.image(0, 0, 'home_background_00').setOrigin(0)
        this.HTMLlogin = this.add.dom(60, 120).createFromCache('login').setOrigin(0)
    }
    update() {}
}
