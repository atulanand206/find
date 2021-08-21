package com.creations.games.think.scenes.main.objects

import com.badlogic.gdx.Input
import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.utils.Align
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.gameObject.addInputListener
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.FontSize
import com.creations.games.think.data.Puzzle
import com.creations.games.think.data.xPuzzles
import com.creations.games.think.data.xTheme
import com.creations.games.think.scenes.BackgroundRounded
import com.creations.games.think.scenes.DarkThemeColor
import com.creations.games.think.utils.assets

class PuzzleFrame(private val di: DI) : GameObject(), DarkThemeColor {

	private lateinit var backgroundRounded: GameObject
	private lateinit var puzzleName: GameObject
	private lateinit var puzzleText: GameObject
	private lateinit var levelSwipe: GameObject
	private var level = 1
	private var isMoving = false
	private var isMovingLeft = false
	private var isMovingRight = false
	private var canMoveLeft = true
	private var canMoveRight = true
	private val puzzles = xPuzzles

	init {
		this.addInputListener()
		addFrameBackground()
		addPuzzle(xTheme.puzzles[level])
		addLevel()
	}

	private fun addFrameBackground() {
		backgroundRounded = BackgroundRounded(di)
		backgroundRounded.setPosition(x, y + Values.VIRTUAL_HEIGHT * 0.48f)
		backgroundRounded.height = Values.VIRTUAL_HEIGHT * 0.42f
		backgroundRounded.width = Values.VIRTUAL_WIDTH
		backgroundRounded.color = viewColor(xTheme.specs)
		addActor(backgroundRounded)
	}

	private fun addPuzzle(puzzle: Puzzle) {
		puzzleName = addPuzzleName(puzzle)
		puzzleText = addPuzzleText(puzzle)
		addActor(puzzleName)
		addActor(puzzleText)
	}

	private fun addLevel() {
		levelSwipe = LevelSwiper(di, level)
		levelSwipe.setPosition(x, y + Values.VIRTUAL_HEIGHT * 0.56f)
		if (level < puzzles.size - 1) canMoveRight = true
		if (level > 0) canMoveLeft = true
		addActor(levelSwipe)
	}

	private fun addPuzzleName(puzzle: Puzzle): GameObject = di.assets.addWrappedLabel(
			puzzle.name, FontSize.F36, viewColor(puzzle.specs),
			x + Values.VIRTUAL_WIDTH * 0.5f, y + Values.VIRTUAL_HEIGHT * 0.85f, Align.center)

	private fun addPuzzleText(puzzle: Puzzle): GameObject = di.assets.addWrappedLabel(
			puzzle.text, FontSize.F18, viewColor(puzzle.specs),
			x + Values.VIRTUAL_WIDTH * 0.5f, y + Values.VIRTUAL_HEIGHT * 0.72f, Align.center)

	override fun touchDragged(x: Float, y: Float, pointer: Int) {
		if (x < 0){
			
		}
		super.touchDragged(x, y, pointer)
	}

	override fun keyDown(keycode: Int): Boolean {
		when (keycode) {
			Input.Keys.LEFT -> {
				isMoving = true
				isMovingLeft = true
			}
			Input.Keys.RIGHT -> {
				isMoving = true
				isMovingRight = true
			}
		}
		return super.keyDown(keycode)
	}

	override fun keyUp(keycode: Int): Boolean {
		when (keycode) {
			Input.Keys.LEFT -> {
				isMoving = false
				isMovingLeft = false
			}
			Input.Keys.RIGHT -> {
				isMoving = false
				isMovingRight = false
			}
		}
		return super.keyUp(keycode)
	}

	override fun act(delta: Float) {
		if (puzzleName.x < -Values.VIRTUAL_WIDTH) {
			puzzleName.x += Values.VIRTUAL_WIDTH
			puzzleText.x += Values.VIRTUAL_WIDTH
			level++
			if (level == puzzles.size) {
				(levelSwipe as LevelSwiper).rightColor(Color.WHITE)
			} else {
//				(levelSwipe as LevelSwiper).rightColor(viewColor(xTheme.puzzles[level].specs))
//				canMoveRight = level < puzzles.size - 1
//				canMoveRight = true
				puzzleName = addPuzzleName(puzzles[level])
				puzzleText = addPuzzleText(puzzles[level])
			}
		}
		if (puzzleName.x > Values.VIRTUAL_WIDTH) {
			puzzleName.x -= Values.VIRTUAL_WIDTH
			puzzleText.x -= Values.VIRTUAL_WIDTH
			level--
			if (level == -1) {
				(levelSwipe as LevelSwiper).leftColor(Color.WHITE)
			} else {
//				(levelSwipe as LevelSwiper).leftColor(viewColor(xTheme.puzzles[level].specs))
//				canMoveLeft = level > 0
//				canMoveLeft = true
				puzzleName = addPuzzleName(puzzles[level])
				puzzleText = addPuzzleText(puzzles[level])
			}
		}

		if (level >= 0 && level < puzzles.size) {
			if (isMoving) {
				if (canMoveLeft && isMovingLeft) {
					puzzleName.moveBy(-2f, 0f)
					puzzleText.moveBy(-2f, 0f)
				}
				if (canMoveRight && isMovingRight) {
					puzzleName.moveBy(2f, 0f)
					puzzleText.moveBy(2f, 0f)
				}
			}
		}
	}
}