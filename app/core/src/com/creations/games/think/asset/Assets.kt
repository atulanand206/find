package com.creations.games.think.asset

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.scenes.scene2d.ui.Label
import com.badlogic.gdx.utils.Align
import com.creations.games.engine.asset.GameAssetManager
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.widgets.TextLabel
import com.creations.games.think.scenes.icons.Icon
import com.creations.games.think.scenes.icons.PlayIcon
import com.creations.games.think.utils.assets

class Assets(private val gam: GameAssetManager) {
	lateinit var circle: Texture
	lateinit var rect: Texture
	lateinit var back: Texture
	lateinit var play: Texture
	lateinit var frame: Texture
	lateinit var light: Texture
	lateinit var dark: Texture
	lateinit var settings: Texture

	init {
		//load assets
		gam.addAssetsToLoad(FileNames.circlePath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.squarePath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.iconBackPath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.iconPlayPath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.iconFramePath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.iconLightPath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.iconDarkPath, Texture::class.java)
		gam.addAssetsToLoad(FileNames.iconSettingsPath, Texture::class.java)

		//load fonts
		FontSize.values().forEach { gam.addFontToLoad(FileNames.fontPath, it.value) }

		gam.addOnLoadListener {
			//automatically added when assets are loaded
			circle = gam.get(FileNames.circlePath, Texture::class.java)
			rect = gam.get(FileNames.squarePath, Texture::class.java)
			back = gam.get(FileNames.iconBackPath, Texture::class.java)
			play = gam.get(FileNames.iconPlayPath, Texture::class.java)
			frame = gam.get(FileNames.iconFramePath, Texture::class.java)
			light = gam.get(FileNames.iconLightPath, Texture::class.java)
			dark = gam.get(FileNames.iconDarkPath, Texture::class.java)
			settings = gam.get(FileNames.iconSettingsPath, Texture::class.java)
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

	fun addLabel(
			text: String,
			size: FontSize,
			color: Color,
			x: Float,
			y: Float,
			align: Int): TextLabel {
		val label = createLabel(text, size, color)
		label.setPosition(x, y, align)
		label.setAlignment(align)
		return label
	}

	fun addWrappedLabel(
			text: String,
			size: FontSize,
			color: Color,
			x: Float,
			y: Float,
			align: Int): GameObject {
		val label = addLabel(text, size, color, x, y, align)
		label.setWrap(true)
		return label
	}

	fun addIcon(icon: Icon, color: Color, x: Float, y: Float, width: Float, height: Float): GameObject {
		icon.color = color
		icon.x = x
		icon.y = y
		icon.width = width
		icon.height = height
		return icon
	}
}