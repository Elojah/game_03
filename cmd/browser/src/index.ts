import { Home } from "cmd/browser/src/scene/home";
import { Login } from "cmd/browser/src/scene/login";
import { Game } from "cmd/browser/src/scene/game";

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
			arcade: { debug: true }
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
