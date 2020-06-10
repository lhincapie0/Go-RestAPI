package com.example.restapi.controller;

import android.animation.ObjectAnimator;
import android.content.Intent;
import android.os.Parcelable;
import android.util.Log;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.Toast;

import com.example.restapi.model.HistoryElement;
import com.example.restapi.util.HTTPSWebUtilDomi;
import com.example.restapi.view.DetailActivity;
import com.example.restapi.view.SearchHistoryActivity;
import com.google.gson.Gson;

import java.util.ArrayList;

import static com.example.restapi.util.Constants.ENDPOINT2;
import static com.example.restapi.util.Constants.INVALID_DOMAIN;
import static com.example.restapi.util.Constants.SEARCH_CALLBACK;

public class SearchHistoryController implements HTTPSWebUtilDomi.OnResponseListener{

    private SearchHistoryActivity activity;
    private HTTPSWebUtilDomi utilDomi;
    private History historyRepo;

    public SearchHistoryController(SearchHistoryActivity activity)
    {
        this.activity = activity;
        utilDomi = new HTTPSWebUtilDomi();
        utilDomi.setListener(this);
        configureLV();
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
                 historyRepo = gson.fromJson(response, History.class);

                ArrayList<String> hosts = new ArrayList<String>();
                int size = historyRepo.items.length;
                for(int i =0; i<size;i++)
                {
                    hosts.add(historyRepo.items[size-1-i].getHost());
                    Log.e("history::: ",historyRepo.getItems()[i].getHost());
                    Log.e("history::: ",historyRepo.getItems()[i].getHost());


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





    public void configureLV()
    {
        activity.getHistoryLV().setClickable(true);
        activity.getHistoryLV().setOnItemClickListener(new AdapterView.OnItemClickListener() {

            @Override
            public void onItemClick(AdapterView<?> arg0, View arg1, int position, long arg3) {
                Object o = activity.getHistoryLV().getItemAtPosition(position);
                String str=(String)o;//As you are using Default String Adapter}
                int size = historyRepo.getItems().length;
                HistoryElement element = historyRepo.getItems()[size-1-position];

                Toast.makeText(activity.getApplicationContext(),element.getHost(),Toast.LENGTH_SHORT).show();
                if(element.getInfo().getTitle().equals("ERROR") || element.getInfo().getTitle().equals(INVALID_DOMAIN))
                {
                    Toast.makeText(activity.getApplicationContext(),"ERROR GETTING INFORMATION",Toast.LENGTH_LONG).show();
                }else
                {
                    Gson gson = new Gson();
                    String info = gson.toJson(element);
                    Intent i = new Intent(activity, DetailActivity.class);
                    i.putExtra("title", element.getInfo().getTitle());
                    i.putExtra("logo", element.getInfo().getLogo());
                    i.putExtra("previous_ssl_grade", element.getInfo().getPrevious_ssl_grade());
                    i.putExtra("ssl_grade", element.getInfo().getSsl_grade());
                    i.putExtra("servers", element.getInfo().getServers());
                    boolean is = element.getInfo().isServers_changed();
                    if(is)
                    {
                        i.putExtra("servers_changed", "True");
                    }else
                    {
                        i.putExtra("servers_changed", "False");
                    }
                    boolean down = element.getInfo().is_down;
                    if(down)
                    {
                        i.putExtra("is_down", "True");
                    }else
                    {
                        i.putExtra("is_down", "False");
                    }
                    activity.startActivity(i);
                }



            }
        });

    }

    public History getHistoryRepo()
    {
        return  historyRepo;
    }
    public static class History {

        private HistoryElement[] items;

        public HistoryElement[] getItems() {
            return items;
        }

        public void setItems(HistoryElement[] items) {
            this.items = items;
        }

        public History() {

        }

        public History(HistoryElement[] items) {
            this.items = items;
        }
    }

}
