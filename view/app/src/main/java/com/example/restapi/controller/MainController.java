package com.example.restapi.controller;

import android.content.Intent;
import android.view.View;

import com.example.restapi.R;
import com.example.restapi.view.DomainSearchActivity;
import com.example.restapi.view.MainActivity;
import com.example.restapi.view.SearchHistoryActivity;

public class MainController implements View.OnClickListener {

    private MainActivity activity;

    public MainController(MainActivity activity){
        this.activity = activity;
        this.activity.getSearchBtn().setOnClickListener(this);
        this.activity.getHistoryBtn().setOnClickListener(this);
    }


    @Override
    public void onClick(View v) {
        switch (v.getId()){
            case R.id.searchBtn:
            {
                Intent i = new Intent(activity, DomainSearchActivity.class);
                activity.startActivity(i);
                break;
            }
            case R.id.historyBtn:
            {
                Intent i = new Intent(activity, SearchHistoryActivity.class);
                activity.startActivity(i);
                break;
            }

        }
    }
}
