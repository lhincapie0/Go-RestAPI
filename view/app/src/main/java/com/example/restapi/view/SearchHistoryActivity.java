package com.example.restapi.view;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.ListView;

import com.example.restapi.R;
import com.example.restapi.controller.SearchHistoryController;

public class SearchHistoryActivity extends AppCompatActivity {

    private ListView historyLV;
    private SearchHistoryController controller;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_search_history);
        historyLV = findViewById(R.id.historyLV);
        this.controller = new SearchHistoryController(this);
    }

    public ListView getHistoryLV() {
        return historyLV;
    }
}
