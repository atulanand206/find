package com.creations.games.think.scenes

import com.badlogic.gdx.graphics.Color
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.scenes.Scene
import com.creations.games.think.data.ViewSpecs

open class GameScene(private val di: DI) : Scene(di) {

	protected var darkTheme = true

	fun viewColor(specs: ViewSpecs): Color {
		return if (darkTheme) specs.dark.color else specs.light.color
	}

	override fun act(dt: Float) {

	}
}