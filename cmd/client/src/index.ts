import { Loading } from "cmd/client/src/scene/loading";
import { Game } from "cmd/client/src/scene/game";

function main() {
	const config: Phaser.Types.Core.GameConfig = {
		title: 'GAME_03',
		type: Phaser.AUTO,
		parent: 'main',
		pixelArt: true,
		scale: {
			parent: 'main',
			mode: Phaser.Scale.FIT,
			autoCenter: Phaser.Scale.CENTER_BOTH,
			min: {
				width: 256,
				height: 144,
			},
			max: {
				width: 1920,
				height: 1080,
			},
			width: window.innerWidth,
			height: window.innerHeight
		},
		dom: {
			createContainer: true
		},
		physics: {
			default: 'arcade',
			arcade: {
				debug: true,
				tileBias: 32
			}
		}
	};

	const g = new Phaser.Game(config)

	g.scene.add('loading', new Loading({
		key: 'loading',
		active: true,
		visible: true,
	}))

	g.scene.add('game', new Game({
		key: 'game',
		active: false,
		visible: false,
	}))

	g.scene.start('loading')
};

main();
