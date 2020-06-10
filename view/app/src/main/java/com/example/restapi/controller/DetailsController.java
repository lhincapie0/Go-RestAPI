package com.example.restapi.controller;

import android.util.Log;

import com.bumptech.glide.Glide;
import com.example.restapi.model.HistoryElement;
import com.example.restapi.model.Server;
import com.example.restapi.view.DetailActivity;
import com.google.gson.Gson;

public class DetailsController {

    private DetailActivity activity;

    public DetailsController(DetailActivity activity)
    {
        this.activity = activity;
         String title = activity.getIntent().getStringExtra("title");
        activity.getTitleTV().setText(title);
       String sslGrade =activity.getIntent().getStringExtra("ssl_grade");
        activity.getSslGradeTV().setText(sslGrade);
        String previousSslGrade = activity.getIntent().getStringExtra("previous_ssl_grade");
        activity.getPrevSslGradeTV().setText(previousSslGrade);
        String serversChanged = activity.getIntent().getStringExtra("servers_changed");
        activity.getServersChangedTV().setText(serversChanged + "");
        String isDown = activity.getIntent().getStringExtra("is_down");
        activity.getsDownTV().setText(isDown);
        String logo = activity.getIntent().getStringExtra("logo");
        Glide.with(activity).load(logo).centerCrop().into(activity.getImageCV());
        String servers = activity.getIntent().getStringExtra("servers");
        Gson gson = new Gson();
    }


}
