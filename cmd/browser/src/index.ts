import {Home} from "@cmd/browser/src/scene/home";
import {Login} from "@cmd/browser/src/scene/login";
import {Game} from "@cmd/browser/src/scene/game";

function main() {

  const config: Phaser.Types.Core.GameConfig = {
    title: 'GAME_03',
    type: Phaser.AUTO,
    parent: 'main',
    pixelArt: true,
    scale: {
      width: window.innerWidth,
      height: window.innerHeight
    },
    dom: {
      createContainer: true
    }
  };

  const g = new Phaser.Game(config)

  g.scene.add('login', new Login({
    key: 'login',
    active: true,
    visible: true,
  }))

  g.scene.add('home', new Home({
    key: 'home',
    active: false,
    visible: false,
  }))

  g.scene.add('game', new Game({
    key: 'game',
    active: false,
    visible: false,
  }))

  g.scene.start('login')
};

main();
