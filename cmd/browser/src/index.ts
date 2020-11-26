import * as Phaser from "phaser";
import {Signin} from "@cmd/browser/src/scene/signin";

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
  g.scene.add('signin', new Signin({
    key: 'signin',
    active: true,
    visible: true,
  }))
  // g.scene.add('lobby', new Lobby({
  //   key: 'lobby',
  //   active: false,
  //   visible: true,
  // }))
  g.scene.start('signin')
};

main();
