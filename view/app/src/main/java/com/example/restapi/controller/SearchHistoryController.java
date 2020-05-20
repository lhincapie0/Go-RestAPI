package com.example.restapi.controller;

import android.util.Log;
import android.widget.ArrayAdapter;

import com.bumptech.glide.Glide;
import com.example.restapi.model.Domain;
import com.example.restapi.util.HTTPSWebUtilDomi;
import com.example.restapi.view.SearchHistoryActivity;
import com.google.gson.Gson;

import java.util.ArrayList;

import static com.example.restapi.util.Constants.ENDPOINT1;
import static com.example.restapi.util.Constants.ENDPOINT2;
import static com.example.restapi.util.Constants.SEARCH_CALLBACK;

public class SearchHistoryController implements HTTPSWebUtilDomi.OnResponseListener {

    private SearchHistoryActivity activity;
    private HTTPSWebUtilDomi utilDomi;

    public SearchHistoryController(SearchHistoryActivity activity)
    {
        this.activity = activity;
        utilDomi = new HTTPSWebUtilDomi();
        utilDomi.setListener(this);
        Log.d("--_>", "SI ESTA IMPRIMIENDO");
        new Thread(
                () ->{
                    utilDomi.GETrequest(SEARCH_CALLBACK,ENDPOINT2);

                }
        ).start();
    }

    @Override
    public void onResponse(int callbackID, String response) {
        switch (callbackID)
        {
            case SEARCH_CALLBACK:
            {
                Log.d("RESPONSE ----------->", response);
                Gson gson = new Gson();
                HostRepo hostRepo = gson.fromJson(response, HostRepo.class);
                ArrayList<String> hosts = new ArrayList<String>();
                int siz = hostRepo.items.length;
                for(int i =0; i<siz;i++)
                {
                    hosts.add(hostRepo.items[siz-1-i]);
                }
                ArrayAdapter<String> itemsAdapter =
                        new ArrayAdapter<String>(activity, android.R.layout.simple_list_item_1, hosts);


                activity.runOnUiThread(
                        () ->
                        {
                            activity.getHistoryLV().setAdapter(itemsAdapter);
                            itemsAdapter.notifyDataSetChanged();
                        }
                );

                break;
            }
        }


    }

    public static class HostRepo{

        private String[] items;

        public String[] getItems() {
            return items;
        }

        public void setItems(String[] items) {
            this.items = items;
        }

        public HostRepo() {

        }

        public HostRepo(String[] items) {
            this.items = items;
        }
    }
}
