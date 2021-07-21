package com.creations.games.think.scenes

import com.badlogic.gdx.graphics.Color
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.scenes.Scene
import com.creations.games.engine.values.Values
import com.creations.games.think.data.ViewSpecs

interface Loader {
	fun load()
	fun unload()
}

interface DarkTheme {
	fun flipDarkMode() {
		Values.isDark = !Values.isDark
	}
}

interface DarkThemeColor {
	fun viewColor(specs: ViewSpecs): Color = if (Values.isDark) specs.dark.color else specs.light.color

	fun viewColorInvert(specs: ViewSpecs): Color = if (Values.isDark) specs.light.color else specs.dark.color
}

abstract class GameScene(private val di: DI) : Scene(di), Loader, DarkThemeColor, DarkTheme {

	override fun flipDarkMode() {
		super.flipDarkMode()
		unload()
		load()
	}

	override fun act(dt: Float) {

	}
}