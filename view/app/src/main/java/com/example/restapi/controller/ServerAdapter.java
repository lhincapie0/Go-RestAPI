package com.example.restapi.controller;

import android.content.res.Resources;
import android.graphics.drawable.Drawable;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import androidx.constraintlayout.widget.ConstraintLayout;
import androidx.core.content.res.ResourcesCompat;

import com.example.restapi.R;
import com.example.restapi.model.Server;

import java.util.ArrayList;


public class ServerAdapter extends BaseAdapter {

    private ArrayList<Server> servers;

    public ServerAdapter()
    {
        servers = new ArrayList<>();
    }

    @Override
    public int getCount() {
        return servers.size();
    }

    @Override
    public Server getItem(int position) {
        return servers.get(position);
    }

    @Override
    public long getItemId(int position) {
        return position;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        LayoutInflater inflater = LayoutInflater.from(parent.getContext());
        View view = inflater.inflate(R.layout.server_line,null);
        TextView ipServerTV = view.findViewById(R.id.ipServerTV);
        ipServerTV.setText("IP Server: "+servers.get(position).getIpAddress());
        TextView sslGradeServerTV = view.findViewById(R.id.sslGradeServerTV);
        sslGradeServerTV.setText("Ssl grade: "+servers.get(position).getGrade());

        TextView countryServerTV = view.findViewById(R.id.countryServerTV);
        countryServerTV.setText("Country: "+servers.get(position).getCountry());

        TextView ownerServerTV = view.findViewById(R.id.ownerServerTV);
        ownerServerTV.setText("Owner: "+servers.get(position).getOwner());


/**
        if(position%2 == 0)
        {
            Drawable drawable = ResourcesCompat.getDrawable(parent.getResources(), R.drawable.bg3, null);
            ConstraintLayout serverLayout = view.findViewById(R.id.serverLayout);
            serverLayout.setBackground(drawable);

        }**/
        return view;

    }




    public void setServers(Server[] serv){
        for(int i = 0; i<serv.length;i++)
        {
            servers.add(serv[i]);
        }

        notifyDataSetChanged();
    }


    public void addServer(Server server)
    {
        servers.add(server);
        this.notifyDataSetChanged();

    }
}