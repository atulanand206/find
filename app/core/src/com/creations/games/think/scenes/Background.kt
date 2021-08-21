package com.creations.games.think.scenes

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.utils.assets

class Background(private val di: DI, private val texture: Texture) : GameObject() {

	override fun draw(batch: Batch, parentAlpha: Float) {
		batch.color = color
		batch.draw(texture, 0f, 0f, Values.VIRTUAL_WIDTH, Values.VIRTUAL_HEIGHT)
		batch.draw(di.assets.imgGray, 0f, 0f, Values.VIRTUAL_WIDTH, Values.VIRTUAL_HEIGHT)
		batch.color = Color.WHITE
		super.draw(batch, parentAlpha)
	}
}