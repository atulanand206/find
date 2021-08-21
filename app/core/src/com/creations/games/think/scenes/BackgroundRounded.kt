package com.creations.games.think.scenes

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.utils.assets


class BackgroundRounded(private val di: DI) : GameObject() {

	override fun draw(batch: Batch, parentAlpha: Float) {
		batch.color = color
		batch.draw(di.assets.imgGray, x, y, width, height)

		batch.color = Color.WHITE
		super.draw(batch, parentAlpha)
	}
}