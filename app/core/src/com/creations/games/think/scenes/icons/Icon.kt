package com.creations.games.think.scenes.icons

import com.badlogic.gdx.Input
import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.utils.logger.Log
import com.creations.games.think.scenes.GameScene
import com.creations.games.think.utils.assets
import com.creations.games.think.utils.scene

abstract class Icon(private val di: DI, private val texture: Texture) : GameObject() {
	abstract fun onClick()

	override fun touchDown(x: Float, y: Float, pointer: Int, button: Int): Boolean {
		if (button != Input.Buttons.LEFT) return false
		val scene = di.scene
		if (scene is GameScene) {
			onClick()
		}
		return super.touchDown(x, y, pointer, button)
	}

	override fun draw(batch: Batch, parentAlpha: Float) {
		batch.color = color
		batch.draw(texture, x - width / 2, y - width / 2, width, height)
		batch.color = Color.WHITE
		super.draw(batch, parentAlpha)
	}
}

class BackIcon(private val di: DI) : Icon(di, di.assets.back) {
	override fun onClick() {
		Log.i("play clicked")
	}
}

class PlayIcon(private val di: DI) : Icon(di, di.assets.play) {
	override fun onClick() {
		Log.i("play clicked")
	}
}

class FrameIcon(private val di: DI) : Icon(di, di.assets.frame) {
	override fun onClick() {
		Log.i("play clicked")
	}
}

class SettingsIcon(private val di: DI) : Icon(di, di.assets.settings) {
	override fun onClick() {
		Log.i("play clicked")
	}
}