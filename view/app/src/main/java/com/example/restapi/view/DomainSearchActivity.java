package com.example.restapi.view;

import androidx.appcompat.app.AppCompatActivity;
import androidx.constraintlayout.widget.ConstraintLayout;

import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.ListView;
import android.widget.TextView;

import com.example.restapi.R;
import com.example.restapi.controller.DomainSearchController;

public class DomainSearchActivity extends AppCompatActivity {

    private DomainSearchController controller;
    private Button searchDomainBtn;
    private EditText domainET;
    private TextView domainTitleTV, domainSslGradeTV, domainIsDownTV, domainPreviousSslGradeTV, domainServersChangedTV, waitingTV;
    private ImageView domainIV;
    private ListView serversLV;
    private ConstraintLayout domainInfoCL;



    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_domain_search);
        searchDomainBtn = findViewById(R.id.searchDomainBtn);
        domainET = findViewById(R.id.domainET);
        domainTitleTV = findViewById(R.id.domainTitleTV);
        domainSslGradeTV = findViewById(R.id.domainSslGradeTV);
        domainPreviousSslGradeTV = findViewById(R.id.domainPreviousSslGradeTV);
        domainServersChangedTV = findViewById(R.id.domainServersChangedTV);
        domainIsDownTV = findViewById(R.id.domainIsDownTV);
        domainIV = findViewById(R.id.domainIV);
        serversLV = findViewById(R.id.serverLV);
        waitingTV = findViewById(R.id.waitingTV);
        domainInfoCL = findViewById(R.id.domainInfoCL);
        controller = new DomainSearchController(this);
    }

    public TextView getWaitingTV() {
        return waitingTV;
    }

    public TextView getDomainTitleTV() {
        return domainTitleTV;
    }

    public TextView getDomainSslGradeTV() {
        return domainSslGradeTV;
    }

    public ImageView getDomainIV() {
        return domainIV;
    }

    public EditText getDomainET() {
        return domainET;
    }

    public Button getSearchDomainBtn() {
        return searchDomainBtn;
    }

    public ListView getServersLV() {
        return serversLV;
    }

    public TextView getDomainIsDownTV() {
        return domainIsDownTV;
    }

    public TextView getDomainPreviousSslGradeTV() {
        return domainPreviousSslGradeTV;
    }

    public TextView getDomainServersChangedTV() {
        return domainServersChangedTV;
    }

    public ConstraintLayout getDomainInfoCL() {
        return domainInfoCL;
    }
}
