package com.example.restapi.controller;

import android.os.AsyncTask;
import android.util.Log;
import android.view.View;
import android.widget.EditText;

import com.bumptech.glide.Glide;
import com.example.restapi.R;
import com.example.restapi.model.Domain;
import com.example.restapi.util.HTTPSWebUtilDomi;
import com.example.restapi.view.DomainSearchActivity;
import com.google.gson.Gson;

import java.io.BufferedInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

import static com.example.restapi.util.Constants.ENDPOINT1;
import static com.example.restapi.util.Constants.MESSAGE_ERROR;
import static com.example.restapi.util.Constants.SEARCH_CALLBACK;

public class DomainSearchController implements View.OnClickListener, HTTPSWebUtilDomi.OnResponseListener{

    private DomainSearchActivity activity;
    private HTTPSWebUtilDomi utilDomi;
    private ServerAdapter serverAdapter;


    public DomainSearchController(DomainSearchActivity activity){
        this.activity = activity;
        this.activity.getSearchDomainBtn().setOnClickListener(this);
        utilDomi = new HTTPSWebUtilDomi();
        utilDomi.setListener(this);
        activity.getDomainInfoCL().setVisibility(View.GONE);



    }

    @Override
    public void onClick(View v) {
        switch(v.getId())
        {
            case R.id.searchDomainBtn:
            {
                activity.getDomainInfoCL().setVisibility(View.GONE);

                String domain = this.activity.getDomainET().getText().toString();
                makeRequest2(domain);
                activity.getSearchDomainBtn().setEnabled(false);
                activity.getWaitingTV().setVisibility(View.VISIBLE);
                activity.getErrorTV().setVisibility(View.GONE);

            }
        }
    }

    public void makeRequest2(String domain){
        new Thread(
                () ->{
                    activity.getDomainInfoCL().setVisibility(View.GONE);

                    Log.d("->>>>>>>>>","GET INFORMATION FOR: "+ domain);
                    Log.e("->>>>>>>>>","GET INFORMATION FOR: "+ domain);
                    utilDomi.GETrequest(SEARCH_CALLBACK,ENDPOINT1+domain);

                }
        ).start();

    }

    @Override
    public void onResponse(int callbackID, String response) {
        switch (callbackID) {
            case SEARCH_CALLBACK: {
                Log.d("RESPONSE ---->",response);
                if(response.equals(MESSAGE_ERROR))
                {
                    activity.runOnUiThread(
                            () ->
                            {
                                activity.getErrorTV().setText(MESSAGE_ERROR);
                                activity.getErrorTV().setVisibility(View.VISIBLE);
                            }
                    );
                }else
                {
                    Gson gson = new Gson();
                    Domain domain = gson.fromJson(response, Domain.class);
                    activity.runOnUiThread(
                            () ->
                            {


                                    activity.getWaitingTV().setVisibility(View.GONE);
                                    activity.getErrorTV().setVisibility(View.GONE);

                                    activity.getDomainInfoCL().setVisibility(View.VISIBLE);
                                    activity.getDomainTitleTV().setText(domain.getTitle());
                                    activity.getDomainSslGradeTV().setText(domain.getSsl_grade());
                                    activity.getDomainPreviousSslGradeTV().setText(domain.getPrevious_ssl_grade());
                                    activity.getDomainIsDownTV().setText(domain.isIs_down()+"");
                                    activity.getDomainServersChangedTV().setText(domain.isServers_changed()+"");
                                    Glide.with(activity).load(domain.getLogo() ).centerCrop().into(activity.getDomainIV());
                                    serverAdapter = new ServerAdapter();
                                    activity.getServersLV().setAdapter(serverAdapter);
                                    serverAdapter.setServers(domain.getServers());
                                    activity.getSearchDomainBtn().setEnabled(true);

                            }
                    );
                }


                break;
            }
        }

    }
}
