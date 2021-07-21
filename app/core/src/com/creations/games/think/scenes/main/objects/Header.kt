package com.creations.games.think.scenes.main.objects

import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.FontSize
import com.creations.games.think.data.Theme
import com.creations.games.think.scenes.DarkThemeColor
import com.creations.games.think.scenes.icons.BackIcon
import com.creations.games.think.utils.assets

class Header(private val di: DI, private val theme: Theme) : GameObject(), DarkThemeColor {

	private var backIcon: GameObject
	private var title: GameObject

	private fun addBackIcon(): GameObject = di.assets.addIcon(BackIcon(di),
			viewColor(theme.specs), x + Values.VIRTUAL_WIDTH * 0.07f, y,
			Values.VIRTUAL_WIDTH * 0.08f, Values.VIRTUAL_HEIGHT * 0.04f)

	private fun addTitle(): GameObject = di.assets.addLabel(
			theme.name, FontSize.F20, color, x + Values.VIRTUAL_WIDTH*0.95f, y, Align.right)

	init {
		backIcon = addBackIcon()
		title = addTitle()
		addActor(backIcon)
		addActor(title)
	}
}