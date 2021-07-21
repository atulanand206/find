package com.creations.games.think.scenes.main.objects

import com.badlogic.gdx.Input
import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.gameObject.addInputListener
import com.creations.games.engine.values.Values
import com.creations.games.think.data.xPuzzle
import com.creations.games.think.scenes.DarkThemeColor
import com.creations.games.think.utils.assets

class ThemeSwitch(private val di: DI) : GameObject(), DarkThemeColor {

	init {
		this.addInputListener()
	}

	override fun keyDown(keycode: Int): Boolean {
		when (keycode) {
			Input.Keys.SPACE -> {
				println("k")
			}
		}
		return super.keyDown(keycode)
	}

	private fun texture() = if (Values.isDark) di.assets.dark else di.assets.light

	override fun draw(batch: Batch, parentAlpha: Float) {
		color = viewColor(xPuzzle.specs)
		batch.draw(texture(), x, y, Values.VIRTUAL_WIDTH * 0.06f, Values.VIRTUAL_HEIGHT * 0.03f)
		color = Color.WHITE
		super.draw(batch, parentAlpha)
	}
}