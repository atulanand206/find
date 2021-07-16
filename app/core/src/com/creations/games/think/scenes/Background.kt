package com.creations.games.think.scenes

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.utils.assets

class Background(private val di: DI): GameObject() {

    private val texture = di.assets.rect

    override fun draw(batch: Batch, parentAlpha: Float) {
        batch.color = color
        batch.draw(texture, x, y, Values.VIRTUAL_WIDTH, Values.VIRTUAL_HEIGHT)

        batch.color = Color.WHITE
        super.draw(batch, parentAlpha)
    }
}