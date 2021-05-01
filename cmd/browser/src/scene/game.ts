import {Scene} from "phaser";
import * as PC from '@pkg/entity/pc_pb';

export class Game extends Scene {

    PC: PC.PC;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    init (pc: PC.PC) {
        this.PC =  pc;
    }
    preload() {}
    create() {}
    update() {}
}
