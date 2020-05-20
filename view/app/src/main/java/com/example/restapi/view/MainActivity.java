package com.example.restapi.view;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.Button;
import android.widget.ImageView;

import com.example.restapi.R;
import com.example.restapi.controller.MainController;

public class MainActivity extends AppCompatActivity {

    private ImageView searchBtn;
    private ImageView historyBtn;
    private MainController controller;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        searchBtn = findViewById(R.id.searchBtn);
        historyBtn = findViewById(R.id.historyBtn);
        controller = new MainController(this);
    }

    public ImageView getSearchBtn() {
        return searchBtn;
    }

    public ImageView getHistoryBtn() {
        return historyBtn;
    }
}
