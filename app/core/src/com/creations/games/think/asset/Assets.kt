package com.creations.games.think.asset

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.scenes.scene2d.ui.Label
import com.creations.games.engine.asset.GameAssetManager
import com.creations.games.engine.widgets.TextLabel

class Assets(private val gam: GameAssetManager) {
    lateinit var circle: Texture
    lateinit var rect: Texture

    var galaxies = ArrayList<Texture>()
    var stars = ArrayList<Texture>()
    lateinit var rocket: Texture
    lateinit var rocketFiring: Texture
    lateinit var nebula: Texture

    lateinit var mine: Texture
    lateinit var flag: Texture
    lateinit var margin: Texture

    init {
        //load assets
        gam.addAssetsToLoad(FileNames.circlePath, Texture::class.java)
        gam.addAssetsToLoad(FileNames.squarePath, Texture::class.java)

        gam.addAssetsToLoad(FileNames.galPath1, Texture::class.java)
        gam.addAssetsToLoad(FileNames.galPath2, Texture::class.java)
        gam.addAssetsToLoad(FileNames.galPath3, Texture::class.java)

        gam.addAssetsToLoad(FileNames.nebulaPath, Texture::class.java)

        gam.addAssetsToLoad(FileNames.starPath1, Texture::class.java)
        gam.addAssetsToLoad(FileNames.starPath2, Texture::class.java)

        gam.addAssetsToLoad(FileNames.rocketPath, Texture::class.java)
        gam.addAssetsToLoad(FileNames.rocketFiringPath, Texture::class.java)

        gam.addAssetsToLoad(FileNames.minePath, Texture::class.java)
        gam.addAssetsToLoad(FileNames.flagPath, Texture::class.java)

        gam.addAssetsToLoad(FileNames.marginPath, Texture::class.java)

        //load fonts
        FontSize.values().forEach { gam.addFontToLoad(FileNames.fontPath, it.value) }

        gam.addOnLoadListener {
            //automatically added when assets are loaded
            circle = gam.get(FileNames.circlePath, Texture::class.java)
            rect = gam.get(FileNames.squarePath, Texture::class.java)

            galaxies.add(gam.get(FileNames.galPath1, Texture::class.java))
            galaxies.add(gam.get(FileNames.galPath2, Texture::class.java))
            galaxies.add(gam.get(FileNames.galPath3, Texture::class.java))

            nebula = gam.get(FileNames.nebulaPath, Texture::class.java)

            stars.add(gam.get(FileNames.starPath1, Texture::class.java))
            stars.add(gam.get(FileNames.starPath2, Texture::class.java))

            rocket = gam.get(FileNames.rocketPath, Texture::class.java)
            rocketFiring = gam.get(FileNames.rocketFiringPath, Texture::class.java)

            mine = gam.get(FileNames.minePath, Texture::class.java)
            flag = gam.get(FileNames.flagPath, Texture::class.java)

            margin = gam.get(FileNames.marginPath, Texture::class.java)
        }
    }

    /**
     * creates a new label with the provided text, color and size
     */
    fun createLabel(
            text: String,
            size: FontSize = FontSize.F12,
            color: Color = Color.WHITE
    ) = TextLabel(text, Label.LabelStyle(gam.getFont(FileNames.fontPath, size.value), color))
}