package com.creations.games.think.scenes.launcher;

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI;
import com.creations.games.engine.scenes.Scene
import com.creations.games.think.asset.FontSize
import com.creations.games.think.asset.Strings
import com.creations.games.think.utils.assets

class SceneLauncher(private val di: DI) : Scene(di) {

	init {
		addTitle()
	}

	private fun addTitle() {
		val title = di.assets.createLabel(Strings.appName, size = FontSize.F48, color = Color.BLACK)
		title.setPosition(screenCenter.x, screenCenter.y, Align.bottom)
		addObjToScene(title)
	}

	override fun act(dt: Float) {
	}
}
