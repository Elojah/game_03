import * as Phaser from "phaser";
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
  g.scene.add('home', new Home({
    key: 'home',
    active: true,
    visible: true,
  }))
  // g.scene.add('lobby', new Lobby({
  //   key: 'lobby',
  //   active: false,
  //   visible: true,
  // }))
  g.scene.start('home')
};

main();
