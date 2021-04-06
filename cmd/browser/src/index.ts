import * as Phaser from "phaser";
import {Login} from "@cmd/browser/src/scene/login";
import {Home} from "@cmd/browser/src/scene/home";

function main() {

  const config: Phaser.Types.Core.GameConfig = {
    title: 'GAME_03',
    type: Phaser.AUTO,
    parent: 'main',
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

  g.scene.start('login')
};

main();
