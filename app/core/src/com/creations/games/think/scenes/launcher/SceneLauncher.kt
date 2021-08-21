package com.creations.games.think.scenes.launcher;

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI;
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.scenes.Scene
import com.creations.games.think.asset.FontSize
import com.creations.games.think.asset.Strings
import com.creations.games.think.data.xTheme
import com.creations.games.think.scenes.Background
import com.creations.games.think.scenes.GameScene
import com.creations.games.think.utils.assets

class SceneLauncher(private val di: DI) : GameScene(di) {

	private lateinit var title: GameObject

	init {
		load()
	}

	private fun addTitle() {
		title = di.assets.createLabel(Strings.appName, size = FontSize.F48, color = Color.BLACK)
		title.setPosition(screenCenter.x, screenCenter.y, Align.bottom)
	}

	override fun load() {
		addTitle()
		addBackground(di.assets.imgLauncher)
		addObjToScene(title)
		addObjToScene(background)
	}

	override fun unload() {
		title.remove()
		background.remove()
	}

	override fun act(dt: Float) {
	}
}
