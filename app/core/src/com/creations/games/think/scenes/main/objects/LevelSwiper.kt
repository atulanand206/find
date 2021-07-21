package com.creations.games.think.scenes.main.objects

import com.badlogic.gdx.Input
import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.gameObject.addInputListener
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.FontSize
import com.creations.games.think.data.Puzzle
import com.creations.games.think.data.xTheme
import com.creations.games.think.scenes.DarkThemeColor
import com.creations.games.think.utils.assets

class LevelSwiper(private val di: DI, private val level: Int): GameObject(), DarkThemeColor {

	private lateinit var levelNo: GameObject
	private lateinit var leftSwipe: GameObject
	private lateinit var rightSwipe: GameObject

	private var isMoving = false

	init {
		this.addInputListener()
		addLevel(xTheme.puzzles[level])
	}

	override fun keyDown(keycode: Int): Boolean {
		when (keycode) {
			Input.Keys.LEFT -> {
				isMoving = true
			}
		}
		return super.keyDown(keycode)
	}

	override fun keyUp(keycode: Int): Boolean {
		when (keycode) {
			Input.Keys.LEFT -> {
				isMoving = false
			}
		}
		return super.keyUp(keycode)
	}

	private fun addLevel(puzzle: Puzzle) {
		levelNo = addLevelNo(puzzle)
		leftSwipe = addLeftSwipeButton(puzzle)
		rightSwipe = addRightSwipeButton(puzzle)
		addActor(levelNo)
		addActor(leftSwipe)
		addActor(rightSwipe)
	}

	private fun addLevelNo(puzzle: Puzzle): GameObject = di.assets.addLabel(
			(level + 1).toString(), FontSize.F100, viewColor(puzzle.specs), x + Values.VIRTUAL_WIDTH*0.5f, y, Align.center)

	private fun addLeftSwipeButton(puzzle: Puzzle): GameObject = di.assets.addLabel(
			"{", FontSize.F36, viewColor(puzzle.specs), x + Values.VIRTUAL_WIDTH*0.07f, y, Align.center)

	private fun addRightSwipeButton(puzzle: Puzzle): GameObject = di.assets.addLabel(
			"}", FontSize.F36, viewColor(puzzle.specs), x + Values.VIRTUAL_WIDTH*0.93f, y, Align.center)

	override fun act(delta: Float) {
		if (isMoving) this.moveBy(10f, 0f);
	}
}