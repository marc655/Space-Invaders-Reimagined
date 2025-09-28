using System;
using System.Drawing;
using System.Windows.Forms;
using System.Collections.Generic;

namespace PacManMazeAdventure
{
    public partial class PacManMazeAdventure : Form
    {
        int[,] maze = {
            {1,1,1,1,1,1,1,1,1,1,1,1},
            {1,0,0,0,1,0,0,0,0,0,0,1},
            {1,0,1,0,1,0,1,1,1,1,0,1},
            {1,0,1,0,0,0,1,0,0,1,0,1},
            {1,0,1,1,1,0,1,0,1,1,0,1},
            {1,0,0,0,1,0,0,0,1,0,0,1},
            {1,1,1,1,1,1,1,1,1,1,1,1}
        };
        int tile = 40, px = 1, py = 1, dir = 0, score = 0;
        List<Point> pellets = new List<Point>();
        Timer timer = new Timer();

        public PacManMazeAdventure()
        {
            InitializeComponent();
            Width = 520; Height = 340;
            KeyPreview = true;
            DoubleBuffered = true;
            for(int y=0;y<maze.GetLength(0);y++)
                for(int x=0;x<maze.GetLength(1);x++)
                    if(maze[y,x]==0) pellets.Add(new Point(x,y));
            timer.Interval = 120;
            timer.Tick += (s,e)=>{ Move(); Invalidate(); };
            timer.Start();
            KeyDown += OnKey;
            Text = "Pac-Man Maze Adventure";
        }

        protected override void OnPaint(PaintEventArgs e)
        {
            base.OnPaint(e);
            for(int y=0;y<maze.GetLength(0);y++)
                for(int x=0;x<maze.GetLength(1);x++)
                    if(maze[y,x]==1)
                        e.Graphics.FillRectangle(Brushes.Blue, x*tile, y*tile, tile, tile);
            foreach(var p in pellets)
                e.Graphics.FillEllipse(Brushes.Yellow, p.X*tile+tile/2-6, p.Y*tile+tile/2-6, 12, 12);
            e.Graphics.FillEllipse(Brushes.Gold, px*tile+8, py*tile+8, tile-16, tile-16);
            e.Graphics.DrawString("Score: "+score, Font, Brushes.White, 440, 12);
        }

        void Move()
        {
            int nx=px, ny=py;
            if(dir==1) nx--; if(dir==2) nx++; if(dir==3) ny--; if(dir==4) ny++;
            if(nx>=0&&ny>=0&&nx<maze.GetLength(1)&&ny<maze.GetLength(0)&&maze[ny,nx]==0)
            { px=nx; py=ny; }
            pellets.RemoveAll(p=>p.X==px&&p.Y==py? (++score>0):false);
        }

        void OnKey(object sender, KeyEventArgs e)
        {
            if(e.KeyCode==Keys.Left) dir=1;
            if(e.KeyCode==Keys.Right) dir=2;
            if(e.KeyCode==Keys.Up) dir=3;
            if(e.KeyCode==Keys.Down) dir=4;
        }

        [STAThread]
        public static void Main()
        {
            Application.EnableVisualStyles();
            Application.Run(new PacManMazeAdventure());
        }
    }
}