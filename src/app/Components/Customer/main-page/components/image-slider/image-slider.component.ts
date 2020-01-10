import { Component, OnInit } from '@angular/core';
import { trigger, transition, animate, style, state } from '@angular/animations';
import { ElementFinder } from 'protractor';

@Component({
  animations: [
    trigger('slideInOut', [
      state('in', style({
        left: '100%'
      })),
      state('out', style({
        left: '-100%'
      })),
      state('inFrame', style({
        left: 0
      })),
      transition('inFrame <=> out', [
        animate('250ms')
      ]),
      transition('inFrame <=> in', [
        animate('250ms')
      ]),
    ])
  ],
  selector: 'app-image-slider',
  templateUrl: './image-slider.component.html',
  styleUrls: ['./image-slider.component.sass']
})
export class ImageSliderComponent implements OnInit {

  currentIndex: number
  doSlide: string
  imageValue: number[] = [1, 2, 3, 4, 5]

  constructor() { }

  ngOnInit() {
    this.doSlide='inFrame'
    this.currentIndex = 1
    setInterval(() => {
      this.changeSlide()
    }, 10000)
  }

  changeSlide() {
    this.doSlide = 'out'
    setTimeout(() => {
      if(this.currentIndex >= 5) this.currentIndex = 1
      else {
        this.currentIndex++
      }
      this.doSlide = 'in'
    }, 250);
    setTimeout(() => {
      this.doSlide = 'inFrame'
    }, 250);
    
  }

  leftClick() {
    this.doSlide = 'in'
    setTimeout(() => {
      if(this.currentIndex >= 5) this.currentIndex = 1
      else {
        this.currentIndex++
      }
      this.doSlide = 'out'
    }, 250);
    setTimeout(() => {
      this.doSlide = 'inFrame'
    }, 250);
  }

  rightClick() {
    this.doSlide = 'out'
    setTimeout(() => {
      if(this.currentIndex >= 5) this.currentIndex = 1
      else {
        this.currentIndex++
      }
      this.doSlide = 'in'
    }, 250);
    setTimeout(() => {
      this.doSlide = 'inFrame'
    }, 250);
  }

  dotsClick(value) {
    if(this.currentIndex > value) {
      this.doSlide = 'in'
      setTimeout(() => {
        this.currentIndex = value
        this.doSlide = 'out'
      }, 250);
      setTimeout(() => {
        this.doSlide = 'inFrame'
      }, 250);
    }
    else if(this.currentIndex < value) {
      this.doSlide = 'out'
      setTimeout(() => {
        this.currentIndex = value
        this.doSlide = 'in'
      }, 250);
      setTimeout(() => {
        this.doSlide = 'inFrame'
      }, 250);
    }
  }

}
