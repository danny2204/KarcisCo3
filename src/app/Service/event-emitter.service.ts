import { Injectable } from '@angular/core';
import { EventEmitter } from '@angular/core';
import { Subscription } from 'rxjs/internal/Subscription'

@Injectable({
  providedIn: 'root'
})
export class EventEmitterService {

  invokeInjectFunction = new EventEmitter()
  invokeAssignFunction = new EventEmitter()

  subsVar: Subscription
  subsVar2: Subscription

  constructor() { }

  onInject(type: string) {
    this.invokeInjectFunction.emit(type)
  }

  onAssign() {
    this.invokeAssignFunction.emit(null)
  }

}
