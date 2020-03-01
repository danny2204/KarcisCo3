import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-history-search',
  templateUrl: './history-search.component.html',
  styleUrls: ['./history-search.component.sass']
})
export class HistorySearchComponent implements OnInit {

  history: Object[] = []

  constructor() { }

  clearHistory() {
    this.history = []
    localStorage.setItem("history", "")
  }

  ngOnInit() {
    if(localStorage.getItem("history") != "") {
      this.history = JSON.parse(localStorage.getItem("history"))
    } else {
      localStorage.setItem("history", "")
    }
  }

}
