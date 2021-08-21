package com.creations.games.think.scenes

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.scenes.Scene
import com.creations.games.engine.values.Values
import com.creations.games.think.data.ViewSpecs
import com.creations.games.think.data.xTheme
import com.creations.games.think.utils.assets

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

	lateinit var background: GameObject

	override fun flipDarkMode() {
		super.flipDarkMode()
		unload()
		load()
	}

	fun addBackground(texture: Texture) {
		background = Background(di, texture)
		addObjToScene(background)
	}

	override fun act(dt: Float) {

	}
}