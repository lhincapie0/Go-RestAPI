
package com.example.restapi.view;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.ListView;
import android.widget.TextView;

import com.example.restapi.R;
import com.example.restapi.controller.DetailsController;

import de.hdodenhof.circleimageview.CircleImageView;

public class DetailActivity extends AppCompatActivity {

    private CircleImageView imageCV;
    private DetailsController controller;
    private TextView titleTV, serversChangedTV, sslGradeTV, prevSslGradeTV, sDownTV;
    private ListView serverslistLV;

    public TextView getsDownTV() {
        return sDownTV;
    }

    public ListView getServerslistLV() {
        return serverslistLV;
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_detail);

        imageCV = findViewById(R.id.imageCV);
        titleTV = findViewById(R.id.titleElementTV);
        serversChangedTV = findViewById(R.id.sChangedTV);
        sslGradeTV = findViewById(R.id.sslGradeTV);
        prevSslGradeTV = findViewById(R.id.prevSslGradeTV);
        serverslistLV = findViewById(R.id.serverslistLV);
        sDownTV = findViewById(R.id.sDownTV);
        this.controller = new DetailsController(this);
    }

    public CircleImageView getImageCV() {
        return imageCV;
    }

    public TextView getTitleTV() {
        return titleTV;
    }

    public TextView getServersChangedTV() {
        return serversChangedTV;
    }

    public TextView getSslGradeTV() {
        return sslGradeTV;
    }

    public TextView getPrevSslGradeTV() {
        return prevSslGradeTV;
    }


}
